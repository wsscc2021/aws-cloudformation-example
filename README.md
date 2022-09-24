
## AWS cloudformation example

This example of cloudformation template for provisioning aws resources. 
The main concept is separate smallest module, regardless structure of various project.
and one module make a independent with anything.

If you using cloudformation, you can write code more detailed configuration than aws-cdk. because, the aws-cdk is high-level constructor that is abstracted from cloudformation.

## Quick Started

1. Clone this repository
    ```
    git clone https://github.com/wsscc2021/aws-cloudformation-example.git .
    ```

2. Modify module, kind of you want provisioning.

3. Create S3 bucket for will upload modified template files.
    - https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html

4. Upload modified template files to s3 bucket which was created.
    - https://docs.aws.amazon.com/AmazonS3/latest/userguide/uploading-an-object-bucket.html

5. Create Cloudformation template
    - https://docs.aws.amazon.com/ko_kr/AWSCloudFormation/latest/UserGuide/using-cfn-cli-creating-stack.html