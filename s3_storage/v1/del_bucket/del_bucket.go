package main

import (
	"fmt"
	"os"
	"strings"

	s3utils "example.com/s3utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Lists all objects in a bucket using pagination
//
// Usage:
// listObjects <bucket>

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("\n\t Error: The bucketName was not provided on the command line")
		lastSlash := strings.LastIndex(os.Args[0], "/")
		programName := os.Args[0]
		programName = programName[lastSlash+1:]
		fmt.Printf("\n\t Try            $> %s bucket-name", programName)
		fmt.Printf("\n\t Instruction: Please provide the bucket name\n\n")
		os.Exit(1)
	}
	sess, err := session.NewSessionWithOptions(
		session.Options{
			Profile: "default",
			Config: aws.Config{
				Region: aws.String("us-east-2"),
			},
		})
	if nil != err {
		fmt.Printf("\n Error Encountered on session initialization:[%v]\n", err)
		return
	}

	s3ClientSvc := s3.New(sess)
	bucketName := os.Args[1]

	bFound, _ := s3utils.DoesBucketExist(bucketName)
	if bFound {
		fmt.Printf("\n The bucket:%s was found \n", bucketName)
	} else {
		fmt.Printf("\n Th bucket:%s was not found \n", bucketName)
		os.Exit(1)
	}

	input := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}

	_, err = s3ClientSvc.DeleteBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

	} else {
		fmt.Printf("\n the bucket: %s was deleted \n\n", bucketName)
	}

}

/*
//
// example.com/storage_utils@v1.0.0-00010101000000-000000000000 (replaced by ../storage_utils)
//
// equire module/path v1.2.3
//

require example.com/s3utils v1.0.0
replace example.com/s3utils => ../s3utils


require github.com/aws/aws-sdk-go v1.46.6
require github.com/jmespath/go-jmespath v0.4.0 // indirect
replace example.com/s3utils => ../s3utils
require example.com/s3utils@v1.0.0-00010101000000-000000000000


*/
