package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
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

	input := &s3.ListBucketsInput{}

	result, err := s3ClientSvc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println(result)
		numOfBuckets := len(result.Buckets)
		for ii := 0; ii < numOfBuckets; ii++ {
			fmt.Printf("\nbucket: %s", *result.Buckets[ii].Name)
		}
	}

} //main
