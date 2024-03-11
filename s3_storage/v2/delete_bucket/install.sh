echo on
go install github.com/aws/aws-sdk-go-v2/config@latest
go install github.com/aws/aws-sdk-go-v2/service/s3@latest
go install github.com/aws/aws-sdk-go/aws@latest
go mod init example.com/delete_buckets_v2
go mod tidy
echo off
