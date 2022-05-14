#!/usr/bin/env python3
""" test fixture deployment
"""
import aws_cdk as cdk

from deployments.deployments_stack import DeploymentsStack


app = cdk.App()
DeploymentsStack(
    app,
    "Postgr8TestDeploymentStack",
)

app.synth()
