#!/usr/bin/env python3
""" test fixture deployment
"""
import aws_cdk as cdk

from deployments.deployments_stack import DeploymentsStack


app = cdk.App()
deployment_stack = DeploymentsStack(
    app,
    "Postgr8TestDeploymentStack",
)
cdk.Tags.of(deployment_stack).add("purpose", "postgr8_test_fixture")
app.synth()
