# kmsca

Utility to generate a (self-signed) CA certificate with a key managed by AWS KMS, and be able to sign CSRs with that CA.

When you generate Certificate Authority (CA) certificate using a private key in AWS Key Management Service (KMS), meaning that the private key never leaves KMS and all signing operations for the certificate will also occur within KMS.

## Step by Step Demo

### (0) Generate the KMS Key

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

### (1) Create a CA certificate from the KMS Key

```
go run __example__/1_create-ca-cert/main.go > ca.pem
```

<details>

<summary>Example</summary>

```
14:56 $ AWS_PROFILE=demo AWS_REGION=us-east-1 go run __example__/1_create-ca-cert/main.go > ca.pem
```

```
14:58 $ cat ca.pem
-----BEGIN CERTIFICATE-----
MIIDVzCCAj+gAwIBAgIIG8CMiEWGiXwwDQYJKoZIhvcNAQELBQAwPzELMAkGA1UE
BhMCQ0ExGjAYBgNVBAoTEUFkcmlhbm8gU2VsYSBJbmMuMRQwEgYDVQQDEwthZHJp
YW5vc2VsYTAeFw0yMzEwMDYyMTUzMzRaFw0zMzEwMDMyMTU4MzRaMD8xCzAJBgNV
BAYTAkNBMRowGAYDVQQKExFBZHJpYW5vIFNlbGEgSW5jLjEUMBIGA1UEAxMLYWRy
aWFub3NlbGEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDC+BDzFcr3
LRM5ITUlqGLXvYWNczxI8lavxTgU5TQPPoS+h6Up99yJzNWzJcjwwDEJdNa0Iffq
ygLYj6Zvbye5hNIXnOKh/4+meFRBAzazgaOq5w6Inl5T0ct1yd9p+oecXZPK27lv
C3BhIx4xUnhrhoH8DkmoiJbyzl52SUWyetu4qMnYA/vVTmvudWuMCYErMAwGAJ7z
IENCi7+DIF/mRNowrDm75yMNNOpWdvbUSF+o9/V83QUPQspkFDP9A8xnAWxJGls5
WsQnDoK2K1k/lpy175sqbgv+rmF4MDYS9zbGyLNaPGJWRrYXQ5lWme03+3WzAEya
5azmjbAP0bEBAgMBAAGjVzBVMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggr
BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTtOdkDNVBgMV8MVNFL
h7IkCSudgjANBgkqhkiG9w0BAQsFAAOCAQEAEhBgslFMVXEb1/fdTmX60rpS2TGY
2VsNr/gEXlEQc1nfODfVEkBkYuRMm9DhZ9rqHo8omruD5MmkzJeR7inatPE5KDBO
5sWbmzTCJjgTC6BtkITG7cVsSFrKtrZ+GO4l2mmB8HQj6SgYf1mebQoRAIb10VdT
CeITJOGYUxo/GmAeceOfEN8CQaJGDZnBniP3zdCLuEgRZQhhYTwIXax4iLE/AHY0
jH9mCfZJ4aM5C3ht1QTN2T23ac7pThS7ZAZtY+lJ6EBgPgE7scD8Sd3zD5rf8CFE
kCq6acbBegVeo4PoFwMbB7+sq7XkJ3KDR7hBaEVWqDKHa0CyxkFdGoEscQ==
-----END CERTIFICATE-----
```

```
14:58 $ cat ca.pem | openssl x509 -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 1999752751462386044 (0x1bc08c884586897c)
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = CA, O = Adriano Sela Inc., CN = adrianosela
        Validity
            Not Before: Oct  6 21:53:34 2023 GMT
            Not After : Oct  3 21:58:34 2033 GMT
        Subject: C = CA, O = Adriano Sela Inc., CN = adrianosela
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:c2:f8:10:f3:15:ca:f7:2d:13:39:21:35:25:a8:
                    62:d7:bd:85:8d:73:3c:48:f2:56:af:c5:38:14:e5:
                    34:0f:3e:84:be:87:a5:29:f7:dc:89:cc:d5:b3:25:
                    c8:f0:c0:31:09:74:d6:b4:21:f7:ea:ca:02:d8:8f:
                    a6:6f:6f:27:b9:84:d2:17:9c:e2:a1:ff:8f:a6:78:
                    54:41:03:36:b3:81:a3:aa:e7:0e:88:9e:5e:53:d1:
                    cb:75:c9:df:69:fa:87:9c:5d:93:ca:db:b9:6f:0b:
                    70:61:23:1e:31:52:78:6b:86:81:fc:0e:49:a8:88:
                    96:f2:ce:5e:76:49:45:b2:7a:db:b8:a8:c9:d8:03:
                    fb:d5:4e:6b:ee:75:6b:8c:09:81:2b:30:0c:06:00:
                    9e:f3:20:43:42:8b:bf:83:20:5f:e6:44:da:30:ac:
                    39:bb:e7:23:0d:34:ea:56:76:f6:d4:48:5f:a8:f7:
                    f5:7c:dd:05:0f:42:ca:64:14:33:fd:03:cc:67:01:
                    6c:49:1a:5b:39:5a:c4:27:0e:82:b6:2b:59:3f:96:
                    9c:b5:ef:9b:2a:6e:0b:fe:ae:61:78:30:36:12:f7:
                    36:c6:c8:b3:5a:3c:62:56:46:b6:17:43:99:56:99:
                    ed:37:fb:75:b3:00:4c:9a:e5:ac:e6:8d:b0:0f:d1:
                    b1:01
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment, Certificate Sign
            X509v3 Extended Key Usage:
                TLS Web Server Authentication
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier:
                ED:39:D9:03:35:50:60:31:5F:0C:54:D1:4B:87:B2:24:09:2B:9D:82
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        12:10:60:b2:51:4c:55:71:1b:d7:f7:dd:4e:65:fa:d2:ba:52:
        d9:31:98:d9:5b:0d:af:f8:04:5e:51:10:73:59:df:38:37:d5:
        12:40:64:62:e4:4c:9b:d0:e1:67:da:ea:1e:8f:28:9a:bb:83:
        e4:c9:a4:cc:97:91:ee:29:da:b4:f1:39:28:30:4e:e6:c5:9b:
        9b:34:c2:26:38:13:0b:a0:6d:90:84:c6:ed:c5:6c:48:5a:ca:
        b6:b6:7e:18:ee:25:da:69:81:f0:74:23:e9:28:18:7f:59:9e:
        6d:0a:11:00:86:f5:d1:57:53:09:e2:13:24:e1:98:53:1a:3f:
        1a:60:1e:71:e3:9f:10:df:02:41:a2:46:0d:99:c1:9e:23:f7:
        cd:d0:8b:b8:48:11:65:08:61:61:3c:08:5d:ac:78:88:b1:3f:
        00:76:34:8c:7f:66:09:f6:49:e1:a3:39:0b:78:6d:d5:04:cd:
        d9:3d:b7:69:ce:e9:4e:14:bb:64:06:6d:63:e9:49:e8:40:60:
        3e:01:3b:b1:c0:fc:49:dd:f3:0f:9a:df:f0:21:44:90:2a:ba:
        69:c6:c1:7a:05:5e:a3:83:e8:17:03:1b:07:bf:ac:ab:b5:e4:
        27:72:83:47:b8:41:68:45:56:a8:32:87:6b:40:b2:c6:41:5d:
        1a:81:2c:71
```

</details>

### (2) Generate a Sample Certificate Signing Request (CSR)

```
go run __example__/2_create-csr/main.go > csr.pem
```

<details>

<summary>Example</summary>

```
14:58 $ AWS_PROFILE=demo AWS_REGION=us-east-1 go run __example__/2_create-csr/main.go > csr.pem
```

```
15:00 $ cat csr.pem
-----BEGIN CERTIFICATE REQUEST-----
MIICgDCCAWgCAQAwOzELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUV4YW1wbGUsIElu
Yy4xFDASBgNVBAMTC2V4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEAwoCDH/AioxszemzK/yXptE6U02cJ3jB/4l7jbCGnp28KFfSgCO31
SaNi7yG1w48p0cW/uOorsO0qUyhL6wt52H9CLmbddNH+/ZWxNT3dT17qhbAc6tRT
vBmLbO9yeTzoYBF5LSRoupndUN7+YsQSu1eEbnsOGjR1xRv9+Ne0q79W3RmaidMU
ywDmyti9drI4Ol788Ebkcf7VTGzZQ1kHlU7CwF5yHZnDe3g98KG+/sdFrNl9EZWm
TpOkE3d0PZv5NltVQdVTpKn/VBXreoOMtErW52AK0RP5mAmhg9wkYI6rFwLBUCSB
gaaSQz3kLZ+07aiTiXOEwn/51IjknnNxoQIDAQABoAAwDQYJKoZIhvcNAQELBQAD
ggEBAE9hVkn/VXdvLYHZlZQNfT2ORjUGnFY5fgoNY6UoqIYyD48p6MttQiFT1Uzo
ZwbkDe3dxCOg2v+hdeYkpCadeJKtsJetF+d5ZSmpXbBa4f/o5mW20F2PeKYEexfQ
eT0c+D0UOQbvtl0iAmWsKlR4Ik0a+8jE5W3v+dYc1XVCB6mdRDzelm5m2MFsiibO
bszN+YGpIf3Ma1jmNvpSy3BCAe5yVaJIFkLJwv3NW+OgDKRoh/K2SIVtuTYUW66l
93eSjPBw9FMaZQE3iM3ffr2+d4sAa18hza2oMom/zsC0a8qKwapCJaC8b9Vnocpv
hC2cRDwvUG2BnxzJ8vOQS45I7XY=
-----END CERTIFICATE REQUEST-----
```

```
15:00 $ cat csr.pem
-----BEGIN CERTIFICATE REQUEST-----
MIICgDCCAWgCAQAwOzELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUV4YW1wbGUsIElu
Yy4xFDASBgNVBAMTC2V4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEAwoCDH/AioxszemzK/yXptE6U02cJ3jB/4l7jbCGnp28KFfSgCO31
SaNi7yG1w48p0cW/uOorsO0qUyhL6wt52H9CLmbddNH+/ZWxNT3dT17qhbAc6tRT
vBmLbO9yeTzoYBF5LSRoupndUN7+YsQSu1eEbnsOGjR1xRv9+Ne0q79W3RmaidMU
ywDmyti9drI4Ol788Ebkcf7VTGzZQ1kHlU7CwF5yHZnDe3g98KG+/sdFrNl9EZWm
TpOkE3d0PZv5NltVQdVTpKn/VBXreoOMtErW52AK0RP5mAmhg9wkYI6rFwLBUCSB
gaaSQz3kLZ+07aiTiXOEwn/51IjknnNxoQIDAQABoAAwDQYJKoZIhvcNAQELBQAD
ggEBAE9hVkn/VXdvLYHZlZQNfT2ORjUGnFY5fgoNY6UoqIYyD48p6MttQiFT1Uzo
ZwbkDe3dxCOg2v+hdeYkpCadeJKtsJetF+d5ZSmpXbBa4f/o5mW20F2PeKYEexfQ
eT0c+D0UOQbvtl0iAmWsKlR4Ik0a+8jE5W3v+dYc1XVCB6mdRDzelm5m2MFsiibO
bszN+YGpIf3Ma1jmNvpSy3BCAe5yVaJIFkLJwv3NW+OgDKRoh/K2SIVtuTYUW66l
93eSjPBw9FMaZQE3iM3ffr2+d4sAa18hza2oMom/zsC0a8qKwapCJaC8b9Vnocpv
hC2cRDwvUG2BnxzJ8vOQS45I7XY=
-----END CERTIFICATE REQUEST-----
```

```
15:00 $ cat csr.pem | openssl req -noout -text
Certificate Request:
    Data:
        Version: 1 (0x0)
        Subject: C = US, O = "Example, Inc.", CN = example.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:c2:80:83:1f:f0:22:a3:1b:33:7a:6c:ca:ff:25:
                    e9:b4:4e:94:d3:67:09:de:30:7f:e2:5e:e3:6c:21:
                    a7:a7:6f:0a:15:f4:a0:08:ed:f5:49:a3:62:ef:21:
                    b5:c3:8f:29:d1:c5:bf:b8:ea:2b:b0:ed:2a:53:28:
                    4b:eb:0b:79:d8:7f:42:2e:66:dd:74:d1:fe:fd:95:
                    b1:35:3d:dd:4f:5e:ea:85:b0:1c:ea:d4:53:bc:19:
                    8b:6c:ef:72:79:3c:e8:60:11:79:2d:24:68:ba:99:
                    dd:50:de:fe:62:c4:12:bb:57:84:6e:7b:0e:1a:34:
                    75:c5:1b:fd:f8:d7:b4:ab:bf:56:dd:19:9a:89:d3:
                    14:cb:00:e6:ca:d8:bd:76:b2:38:3a:5e:fc:f0:46:
                    e4:71:fe:d5:4c:6c:d9:43:59:07:95:4e:c2:c0:5e:
                    72:1d:99:c3:7b:78:3d:f0:a1:be:fe:c7:45:ac:d9:
                    7d:11:95:a6:4e:93:a4:13:77:74:3d:9b:f9:36:5b:
                    55:41:d5:53:a4:a9:ff:54:15:eb:7a:83:8c:b4:4a:
                    d6:e7:60:0a:d1:13:f9:98:09:a1:83:dc:24:60:8e:
                    ab:17:02:c1:50:24:81:81:a6:92:43:3d:e4:2d:9f:
                    b4:ed:a8:93:89:73:84:c2:7f:f9:d4:88:e4:9e:73:
                    71:a1
                Exponent: 65537 (0x10001)
        Attributes:
            (none)
            Requested Extensions:
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        4f:61:56:49:ff:55:77:6f:2d:81:d9:95:94:0d:7d:3d:8e:46:
        35:06:9c:56:39:7e:0a:0d:63:a5:28:a8:86:32:0f:8f:29:e8:
        cb:6d:42:21:53:d5:4c:e8:67:06:e4:0d:ed:dd:c4:23:a0:da:
        ff:a1:75:e6:24:a4:26:9d:78:92:ad:b0:97:ad:17:e7:79:65:
        29:a9:5d:b0:5a:e1:ff:e8:e6:65:b6:d0:5d:8f:78:a6:04:7b:
        17:d0:79:3d:1c:f8:3d:14:39:06:ef:b6:5d:22:02:65:ac:2a:
        54:78:22:4d:1a:fb:c8:c4:e5:6d:ef:f9:d6:1c:d5:75:42:07:
        a9:9d:44:3c:de:96:6e:66:d8:c1:6c:8a:26:ce:6e:cc:cd:f9:
        81:a9:21:fd:cc:6b:58:e6:36:fa:52:cb:70:42:01:ee:72:55:
        a2:48:16:42:c9:c2:fd:cd:5b:e3:a0:0c:a4:68:87:f2:b6:48:
        85:6d:b9:36:14:5b:ae:a5:f7:77:92:8c:f0:70:f4:53:1a:65:
        01:37:88:cd:df:7e:bd:be:77:8b:00:6b:5f:21:cd:ad:a8:32:
        89:bf:ce:c0:b4:6b:ca:8a:c1:aa:42:25:a0:bc:6f:d5:67:a1:
        ca:6f:84:2d:9c:44:3c:2f:50:6d:81:9f:1c:c9:f2:f3:90:4b:
        8e:48:ed:76
```

</details>

### (3) Sign the Sample Certificate Signing Request (CSR) with the CA Cert

```
go run __example__/3_sign-csr/main.go ./ca.pem ./csr.pem > signed-cert.pem
```

<details>

<summary>Example</summary>

```
15:02 $ AWS_PROFILE=demo AWS_REGION=us-east-1 go run __example__/3_sign-csr/main.go ./ca.pem ./csr.pem > signed-cert.pem
```

```
15:02 $ cat signed-cert.pem
-----BEGIN CERTIFICATE-----
MIIDLzCCAhegAwIBAgIId2XxFuBNKJcwDQYJKoZIhvcNAQELBQAwPzELMAkGA1UE
BhMCQ0ExGjAYBgNVBAoTEUFkcmlhbm8gU2VsYSBJbmMuMRQwEgYDVQQDEwthZHJp
YW5vc2VsYTAeFw0yMzEwMDYyMTU3MzBaFw0yNDEwMDUyMjAyMzBaMDsxCzAJBgNV
BAYTAlVTMRYwFAYDVQQKEw1FeGFtcGxlLCBJbmMuMRQwEgYDVQQDEwtleGFtcGxl
LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMKAgx/wIqMbM3ps
yv8l6bROlNNnCd4wf+Je42whp6dvChX0oAjt9UmjYu8htcOPKdHFv7jqK7DtKlMo
S+sLedh/Qi5m3XTR/v2VsTU93U9e6oWwHOrUU7wZi2zvcnk86GAReS0kaLqZ3VDe
/mLEErtXhG57Dho0dcUb/fjXtKu/Vt0ZmonTFMsA5srYvXayODpe/PBG5HH+1Uxs
2UNZB5VOwsBech2Zw3t4PfChvv7HRazZfRGVpk6TpBN3dD2b+TZbVUHVU6Sp/1QV
63qDjLRK1udgCtET+ZgJoYPcJGCOqxcCwVAkgYGmkkM95C2ftO2ok4lzhMJ/+dSI
5J5zcaECAwEAAaMzMDEwDgYDVR0PAQH/BAQDAgWgMB8GA1UdIwQYMBaAFO052QM1
UGAxXwxU0UuHsiQJK52CMA0GCSqGSIb3DQEBCwUAA4IBAQAYsmREf2tDCsKHWqPm
IigXGHKXT5Rbtvkx9FRluj7T3LMAPufc4qZz1fK94qoN2mYcYDuMgOwivkHhfWkk
N9pShoZ+qmbqi3/qqFlq6vPV/xhBrWN0sXl7DIjnhr6zxLR+5IzVgorrqudRRKq5
AfcAwBzhHtiMhiXI6ChuLGnItG6C1P4ZC3hp36NB5Wn6moeL6Y+wZasgiZ/88z72
FN6C9bUD2zv/5I/747YaqNRWt+HmIdSyI42gX7MNSXoGAQk1OfQHX/o8hBDUjGtS
WvBQ1TXcFi4Kb1HbL8Eh8qmkiAuwnvRJSP3v1ZWLs2lQfXMXwBMhgxgGWYemetaJ
fdEc
-----END CERTIFICATE-----
```

```
15:02 $ cat signed-cert.pem | openssl x509 -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 8603547743715928215 (0x7765f116e04d2897)
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = CA, O = Adriano Sela Inc., CN = adrianosela
        Validity
            Not Before: Oct  6 21:57:30 2023 GMT
            Not After : Oct  5 22:02:30 2024 GMT
        Subject: C = US, O = "Example, Inc.", CN = example.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:c2:80:83:1f:f0:22:a3:1b:33:7a:6c:ca:ff:25:
                    e9:b4:4e:94:d3:67:09:de:30:7f:e2:5e:e3:6c:21:
                    a7:a7:6f:0a:15:f4:a0:08:ed:f5:49:a3:62:ef:21:
                    b5:c3:8f:29:d1:c5:bf:b8:ea:2b:b0:ed:2a:53:28:
                    4b:eb:0b:79:d8:7f:42:2e:66:dd:74:d1:fe:fd:95:
                    b1:35:3d:dd:4f:5e:ea:85:b0:1c:ea:d4:53:bc:19:
                    8b:6c:ef:72:79:3c:e8:60:11:79:2d:24:68:ba:99:
                    dd:50:de:fe:62:c4:12:bb:57:84:6e:7b:0e:1a:34:
                    75:c5:1b:fd:f8:d7:b4:ab:bf:56:dd:19:9a:89:d3:
                    14:cb:00:e6:ca:d8:bd:76:b2:38:3a:5e:fc:f0:46:
                    e4:71:fe:d5:4c:6c:d9:43:59:07:95:4e:c2:c0:5e:
                    72:1d:99:c3:7b:78:3d:f0:a1:be:fe:c7:45:ac:d9:
                    7d:11:95:a6:4e:93:a4:13:77:74:3d:9b:f9:36:5b:
                    55:41:d5:53:a4:a9:ff:54:15:eb:7a:83:8c:b4:4a:
                    d6:e7:60:0a:d1:13:f9:98:09:a1:83:dc:24:60:8e:
                    ab:17:02:c1:50:24:81:81:a6:92:43:3d:e4:2d:9f:
                    b4:ed:a8:93:89:73:84:c2:7f:f9:d4:88:e4:9e:73:
                    71:a1
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Authority Key Identifier:
                ED:39:D9:03:35:50:60:31:5F:0C:54:D1:4B:87:B2:24:09:2B:9D:82
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        18:b2:64:44:7f:6b:43:0a:c2:87:5a:a3:e6:22:28:17:18:72:
        97:4f:94:5b:b6:f9:31:f4:54:65:ba:3e:d3:dc:b3:00:3e:e7:
        dc:e2:a6:73:d5:f2:bd:e2:aa:0d:da:66:1c:60:3b:8c:80:ec:
        22:be:41:e1:7d:69:24:37:da:52:86:86:7e:aa:66:ea:8b:7f:
        ea:a8:59:6a:ea:f3:d5:ff:18:41:ad:63:74:b1:79:7b:0c:88:
        e7:86:be:b3:c4:b4:7e:e4:8c:d5:82:8a:eb:aa:e7:51:44:aa:
        b9:01:f7:00:c0:1c:e1:1e:d8:8c:86:25:c8:e8:28:6e:2c:69:
        c8:b4:6e:82:d4:fe:19:0b:78:69:df:a3:41:e5:69:fa:9a:87:
        8b:e9:8f:b0:65:ab:20:89:9f:fc:f3:3e:f6:14:de:82:f5:b5:
        03:db:3b:ff:e4:8f:fb:e3:b6:1a:a8:d4:56:b7:e1:e6:21:d4:
        b2:23:8d:a0:5f:b3:0d:49:7a:06:01:09:35:39:f4:07:5f:fa:
        3c:84:10:d4:8c:6b:52:5a:f0:50:d5:35:dc:16:2e:0a:6f:51:
        db:2f:c1:21:f2:a9:a4:88:0b:b0:9e:f4:49:48:fd:ef:d5:95:
        8b:b3:69:50:7d:73:17:c0:13:21:83:18:06:59:87:a6:7a:d6:
        89:7d:d1:1c
```

</details>