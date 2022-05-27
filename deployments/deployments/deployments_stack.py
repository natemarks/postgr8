#!/usr/bin/env python3
""" RDS Postgres deployment stack
Create a standard vpc and put rds on public subnets
"""
from aws_cdk import (
    # Duration,
    Stack,
    aws_ec2 as ec2,
    aws_rds as rds,
    aws_secretsmanager as secretsmanager
)
from constructs import Construct
import json


class DeploymentsStack(Stack):
    """RDS Postgres Deployment Stack"""

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # create the RDS VPC
        self.db_vpc = ec2.Vpc(
            self,
            "postgr8_test_deployment_stack",
            max_azs=2,
            cidr="10.7.0.0/16",
            vpc_name="postgr8_test_deployment_stack",
        )

        # create the DB SG
        self.db_sg = ec2.SecurityGroup(
            self,
            "postgr8_test_sg",
            vpc=self.db_vpc,
            allow_all_outbound=True,
        )
        # permit inbound access from anywhere
        self.db_sg.add_ingress_rule(
            ec2.Peer.any_ipv4(), ec2.Port.tcp(5432), "PUBLIC POSTGRES ACCESS"
        )

        # create the master secret
        self.secret = secretsmanager.Secret(self,"postgr8_test_secret",
            generate_secret_string=secretsmanager.SecretStringGenerator(
                secret_string_template=json.dumps({"username": "postgres"}),
                generate_string_key="password"
            )
)



        self.cluster = rds.DatabaseInstance(
            self,
            "Postgr8Test",
            engine=rds.DatabaseInstanceEngine.postgres(
                version=rds.PostgresEngineVersion.VER_13_6
            ),
            credentials=rds.Credentials.from_secret(self.secret),
            instance_type=ec2.InstanceType.of(
                ec2.InstanceClass.BURSTABLE3, ec2.InstanceSize.SMALL
            ),
            security_groups=[self.db_sg],
            vpc_subnets=ec2.SubnetSelection(subnet_type=ec2.SubnetType.PUBLIC),
            vpc=self.db_vpc,
        )
