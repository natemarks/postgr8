#!/usr/bin/env python3
""" Test
"""
import aws_cdk as core

from deployments.deployments_stack import DeploymentsStack


def test_sqs_queue_created():
    """test VPC and RDS created"""
    app = core.App()
    stack = DeploymentsStack(app, "deployments")
    template = core.assertions.Template.from_stack(stack)

    template.has_resource_properties("AWS::SQS::Queue", {"VisibilityTimeout": 300})
