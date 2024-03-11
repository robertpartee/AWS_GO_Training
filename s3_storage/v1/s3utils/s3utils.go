package s3utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DoesBucketExist(bucketName string) (bBucketFound bool, errFound error) {
	bBucketFound = false
	errFound = nil

	sess, err := session.NewSessionWithOptions(
		session.Options{
			Profile: "default",
			Config: aws.Config{
				Region: aws.String("us-east-2"),
			},
		})
	if nil != err {
		fmt.Printf("\n Error Encountered on session initialization:[%v]\n", err)
		errFound = err
	}

	svc := s3.New(sess)
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
		errFound = err
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

		nBuckets := len(result.Buckets)
		bBucketFound = false
		for ii := 0; ii < nBuckets; ii++ {
			if *result.Buckets[ii].Name == bucketName {
				bBucketFound = true
				//fmt.Printf("\nThe Bucket Named: %s exisits", *result.Buckets[ii].Name)
			}
		}
	}
	return bBucketFound, errFound
}

// ListObjectsOutput
func ListBucketContents(svc *s3.S3, bucketName string) (listObjectsOutput *s3.ListObjectsOutput, listErr error) {

	input := &s3.ListObjectsInput{
		Bucket:  aws.String(bucketName),
		MaxKeys: aws.Int64(2),
	}

	listObjectsOutput, listErr = svc.ListObjects(input)
	if listErr != nil {
		if aerr, ok := listErr.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(listErr.Error())
		}
	} else {
		//fmt.Println(listObjectsOutput)

	}
	return listObjectsOutput, listErr
}

// example.com/storage_utils@v1.0.0-00010101000000-000000000000
// require module/path v1.2.3
