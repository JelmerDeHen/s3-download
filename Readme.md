EC2-S3-Download
===

About
==

A simple tool to download objects from S3 on EC2 instances. It was written to retrieve private files via IAM rules for instances when bootstrapping machines. Written in go it can be used as an easily deployable alternative to other tools on reduced environments like CoreOS.

Usage
==

Arguments can be substituted with env variables (arguments prefixed by `S3_`, dashes are replaced by underscores).

```
./ec2-s3-download -bucket mybucket -file-path somefile.txt -dest-path=/opt/somefile.txt -region eu-central-1
S3_REGION=eu-central-1 ./ec2-s3-download -bucket mybucket -file-path somefile.txt -dest-path=/opt/somefile.txt
```

Build
==

Go 1.5+ on OSX targetting linux:

```
go get -u github.com/aws/aws-sdk-go
GOOS=linux go build ec2-s3-download.go
```
