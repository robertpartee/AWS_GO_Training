package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

//++
//		Execute the following commands to import the required packages
//
//		go mod init example.com/delete_buckets_v2
//		go install github.com/aws/aws-sdk-go-v2/config@latest
//		go install github.com/aws/aws-sdk-go-v2/service/s3@latest
//		go install github.com/aws/aws-sdk-go/aws@latest
//
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
	bucketName := os.Args[1]

	_, err = s3Client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName)})

	if err != nil {
		fmt.Printf("\nCouldn't delete bucket %v. Here's why: %v\n", bucketName, err)
		os.Exit(1)
	}

	fmt.Printf("\n Bucket:%v  has been deleted ", bucketName)
	fmt.Printf("\n Done \n")
}
