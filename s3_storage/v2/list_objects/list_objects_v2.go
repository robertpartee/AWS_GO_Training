package main

import (
	"context"
	"fmt"
	"log"

	"os"
	"strings"

	// /Users/robpartee/workspace/aws_cloud_class_golang/golang/s3_storage/v1

	"example.com/s3utils_v2"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//++
//		Execute the following commands to import the required packages
//
//		go install github.com/aws/aws-sdk-go-v2/config@latest
//		go install github.com/aws/aws-sdk-go-v2/service/s3@latest
//		go mod init example.com/list_buckets_v2
//		go mod tidy
//
//--

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

	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	bucketName := os.Args[1]

	bFound, _ := s3utils_v2.DoesBucketExist(bucketName)
	if bFound {
		fmt.Printf("\nThe bucket:%s was found \n", bucketName)
	}

	// Set optional parameters (e.g., prefix, delimiter)
	params := &s3.ListObjectsV2Input{
		Bucket: &bucketName,
		// Prefix:    aws.String("your-prefix"), // Uncomment and set if you want to filter by prefix
		// Delimiter: aws.String("/"),           // Uncomment and set if you want to filter by delimiter
	}

	// List objects in the bucket
	resp, err := s3Client.ListObjectsV2(context.TODO(), params)
	if err != nil {
		log.Fatal("Error listing objects:", err)
	}

	// Print object keys
	for _, obj := range resp.Contents {
		fmt.Println("Object key:", *obj.Key)
	}

}

/*



	input := &s3Client.ListObjectsInput{
		Bucket:  aws.String(bucketName),
		MaxKeys: aws.Int64(2),
	}

	result, err := s3Client.ListObjects(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	} else {

		numOfObjects := len(result.Contents)
		if 0 == numOfObjects {
			fmt.Printf("There are no objects in this bucket\n")
		} else {
			fmt.Printf("\nNumber of items:%d\n", numOfObjects)
			for ii := 0; ii < len(result.Contents); ii++ {
				fmt.Printf("Key: %v\n", *result.Contents[ii].Key)
			}
		}
		fmt.Printf("\n %v \n\n", result)
	}



*/
