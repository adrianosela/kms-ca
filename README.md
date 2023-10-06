# kmsca

Utility to generate a (self-signed) CA certificate with a key managed by AWS KMS, and be able to sign CSRs with that CA.

When you generate Certificate Authority (CA) certificate using a private key in AWS Key Management Service (KMS), meaning that the private key never leaves KMS and all signing operations for the certificate will also occur within KMS.

### Generate the KMS Key

There is sample cloudformation yaml in [`./assets/kms-key-cfn.yaml`](./assets/kms-key-cfn.yaml) which you can use to create the key.

You can apply it to your AWS environment with the AWS CLI:

```
aws cloudformation create-stack --stack-name "${YOUR_STACK_NAME}" --template-body "$(cat ./assets/kms-key-cfn.yaml)"
```

Example:

```
10:25 $ AWS_PROFILE=demo AWS_REGION=us-east-1 aws cloudformation create-stack --stack-name kms-ca-key --template-body "$(cat ./assets/kms-key-cfn.yaml)"
{
    "StackId": "arn:aws:cloudformation:us-east-1:123456789012:stack/kms-ca-key/606504b0-646d-11ee-a76a-0a25da907257"
}
```

Alternatively, you can navigate to the AWS Console > KMS Service > Choose "Create Key" and follow the prompts to create a new symmetric of asymmetric key (choose RSA asymmetric operations like signing).