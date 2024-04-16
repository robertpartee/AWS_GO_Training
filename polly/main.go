package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"

	"fmt"
	"os"
)

func main() {
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create Polly client
	svc := polly.New(sess)

	// Get voices for US English
	input := &polly.DescribeVoicesInput{LanguageCode: aws.String("en-US")}

	resp, err := svc.DescribeVoices(input)
	if err != nil {
		fmt.Println("Got error calling DescribeVoices:")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	for _, v := range resp.Voices {
		fmt.Printf("\n\tName:%s\tGender:%v", *v.Name, *v.Gender)
	}
	fmt.Printf("\n\n")
}
