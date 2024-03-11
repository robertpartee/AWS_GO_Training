package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
)

//++
//		Execute the following commands to import the required packages
//
//		go install github.com/aws/aws-sdk-go-v2/config@latest
//		go install github.com/aws/aws-sdk-go-v2/service/s3@latest
//		go mod init example.com/make_buckets_v2
//		go mod tidy
//
//--

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("\n\t Error: The bucketName was not provided on the command line\n")
		lastSlash := strings.LastIndex(os.Args[0], "/")
		programName := os.Args[0]
		programName = programName[lastSlash+1:]
		fmt.Printf("\n\t Try            $> %s bucket-name", programName)
		fmt.Printf("\n\t Instruction: Please provide the bucket name\n\n")
		os.Exit(1)
	}

	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		os.Exit(1)
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	region := "us-east-2"
	bucketName := os.Args[1]

	result, err := s3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})
	if err != nil {
		fmt.Printf("Error encountered: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n Result:%v \n", *result.Location)
	fmt.Printf("\n Done \n")
}
