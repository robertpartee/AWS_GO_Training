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

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("\n\t Error: The bucketName was not provided on the command line")
		lastSlash := strings.LastIndex(os.Args[0], "/")
		programName := os.Args[0]
		programName = programName[lastSlash+1:]
		fmt.Printf("\n\t Try            $> %s bucket-name file-name", programName)
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
	ObjectName := os.Args[2]

	bFound, _ := s3utils.DoesBucketExist(bucketName)
	if bFound {
		fmt.Printf("\n The bucket:%s was found \n", bucketName)
	} else {
		fmt.Printf("\n Th bucket:%s was not found \n", bucketName)
		os.Exit(1)
	}

	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader(ObjectName)),
		Bucket: aws.String(bucketName),
		Key:    aws.String("Christ1"),
	}

	result, err := s3ClientSvc.PutObject(input)
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
			os.Exit(1)
		}

	} else {
		fmt.Println(result)
	}

} //main
