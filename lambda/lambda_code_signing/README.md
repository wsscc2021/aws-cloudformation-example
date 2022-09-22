## Background

Cloudformation does not support signing job.

You can't sign your code for deploy function through cloudformation.

So, You have to split the work that is deploying function from signed code.

## Quick Started

1. Create S3 Bucket

2. Upload `index.zip` (your code)

3. Deploy `lambda_code_signing_1.yaml`, It will be created aws signer profile.

4. Start Signing Job on `AWS Signer` from `index.zip` (your code), It will be created signed code.

5. Deploy `lambda_code_signing_2.yaml`, It will be created lambda function from signed code.