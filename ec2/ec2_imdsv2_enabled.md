
## Background

The cloudformation does not support IMDSv2 options on `AWS::EC2::Instance`
resource type, so in this case, we have to turn using aws-cli.

## Execution

enbale to IMDSv2
```
aws ec2 modify-instance-metadata-options \
    --instance-id i-1234567898abcdef0 \
    --http-tokens required \
    --http-endpoint enabled
```

back to IMDSv1
```
aws ec2 modify-instance-metadata-options \
    --instance-id i-1234567898abcdef0 \
    --http-tokens optional \
    --http-endpoint enabled
```

disable IMDS
```
aws ec2 modify-instance-metadata-options \
    --instance-id i-1234567898abcdef0 \
    --http-endpoint disabled
```