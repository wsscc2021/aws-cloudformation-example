#!/bin/bash
clustername=example-ecs-cluster
AWS_REGION=us-east-1

aws cloudformation create-stack \
    --stack-name CWAgentECS-$clustername-${AWS_REGION} \
    --template-body "$(curl -Ls https://raw.githubusercontent.com/aws-samples/amazon-cloudwatch-container-insights/latest/ecs-task-definition-templates/deployment-mode/daemon-service/cwagent-ecs-instance-metric/cloudformation-quickstart/cwagent-ecs-instance-metric-cfn.json)" \
    --parameters ParameterKey=ClusterName,ParameterValue=$clustername ParameterKey=CreateIAMRoles,ParameterValue=True \
    --capabilities CAPABILITY_NAMED_IAM \
    --region ${AWS_REGION}