package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
)

var bucket string
var objectPath string
var destPath string
var region string

func retrieveFile(key string, bucket string, region string, destPath string) error {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(region)}))
	params := &s3.GetObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)}
	res, err := svc.GetObject(params)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if destPath == "" {
		io.Copy(os.Stdout, res.Body)
		return nil
	}

	outFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	io.Copy(outFile, res.Body)

	return nil
}

func main() {
	flag.StringVar(&bucket, "bucket", os.Getenv("S3_BUCKET"), "s3 bucket")
	flag.StringVar(&region, "region", os.Getenv("S3_REGION"), "aws region")
	flag.StringVar(&objectPath, "object-path", os.Getenv("S3_OBJECT_PATH"), "object path (w/o bucket)")
	flag.StringVar(&destPath, "dest-path", os.Getenv("S3_DEST_PATH"), "destination path (optional)")
	flag.Parse()

	if bucket == "" || objectPath == "" || region == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := retrieveFile(objectPath, bucket, region, destPath)
	if err != nil {
		panic(err)
	}
}
