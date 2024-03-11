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
//		go mod init example.com/put_object_v2
//		go install github.com/aws/aws-sdk-go-v2/config@latest
//		go install github.com/aws/aws-sdk-go-v2/service/s3@latest
//		go install github.com/aws/aws-sdk-go/aws@latest
//
//		go mod tidy
//
//--

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("\n\t Error: The bucketName was not provided on the command line\n")
		lastSlash := strings.LastIndex(os.Args[0], "/")
		programName := os.Args[0]
		programName = programName[lastSlash+1:]
		fmt.Printf("\n\t Try       $> %s bucket-name object-name file-name", programName)
		fmt.Printf("\n\t Instruction: Please provide the bucket name, the object-name, file-name\n\n")
		os.Exit(1)
	}

	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	bucketName := os.Args[1]
	objectKey := os.Args[2]
	fileName := os.Args[3]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
			fileName, bucketName, objectKey, err)
		os.Exit(1)
	}

	fmt.Printf("\n Done \n")
}
