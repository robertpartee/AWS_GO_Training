echo on
go install github.com/aws/aws-sdk-go-v2/config@latest
go install github.com/aws/aws-sdk-go-v2/service/s3@latest
go mod init example.com/list_buckets_v2
go mod tidy
echo off
