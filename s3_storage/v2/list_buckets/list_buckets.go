package main

import (
	"context"
	"fmt"

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
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
		return
	}
	if len(result.Buckets) == 0 {
		fmt.Println("You don't have any buckets!")
	} else {
		//++
		//	shown below are 2 standard methods to iterate through
		//	an array/list of bucket names.  Here is method #1
		//--
		count := len(result.Buckets)
		for _, bucket := range result.Buckets[:count] {
			fmt.Printf("\t%v\n", *bucket.Name)
		}
		//++
		//	shown below is Method #2 which is another example on how to
		//	iterate through the array/list of bucket names
		//--
		fmt.Printf("\n================================================\n")
		for ii := 0; ii < count; ii++ {
			fmt.Printf("\t%2v %v\n", ii, *result.Buckets[ii].Name)
		}
	}
}
