#!/usr/bin/env python3
""" RDS Postgres deployment stack
Create a standard vpc and put rds on public subnets
"""
from aws_cdk import (
    # Duration,
    Stack,
    aws_ec2 as ec2,
    aws_rds as rds,
)
from constructs import Construct


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

        self.db_sg = ec2.SecurityGroup(
            self,
            "LaunchTemplateSG",
            vpc=self.db_vpc,
            allow_all_outbound=True,
        )
        self.db_sg.add_ingress_rule(
            ec2.Peer.any_ipv4(), ec2.Port.tcp(5432), "PUBLIC POSTGRES ACCESS"
        )

        self.cluster = rds.DatabaseInstance(
            self,
            "Postgr8Test",
            engine=rds.DatabaseInstanceEngine.postgres(
                version=rds.PostgresEngineVersion.VER_13_6
            ),
            credentials=rds.Credentials.from_generated_secret(
                "clusteradmin"
            ),  # Optional - will default to 'admin' username and generated password
            instance_type=ec2.InstanceType.of(
                ec2.InstanceClass.BURSTABLE3, ec2.InstanceSize.SMALL
            ),
            security_groups=[self.db_sg],
            vpc_subnets=ec2.SubnetSelection(subnet_type=ec2.SubnetType.PUBLIC),
            vpc=self.db_vpc,
        )
