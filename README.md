
# AWS-CloudFormation-Example

This example of cloudformation template can write code more detailed configuration than aws-cdk.

> The aws-cdk is abstracted to a high-level constructor, it's can simply write code.
> but, it can not support more detailed configuration such as Name tag.

## Work Flow

1. write cloudformation stacks at yaml.

2. create bucket for will upload template.

3. upload template to bucket.

4. create cloudformation stack using template.