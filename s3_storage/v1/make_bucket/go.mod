module make_bucket

go 1.21.1

replace example.com/s3utils => ../s3utils

require (
	example.com/s3utils v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go v1.48.3
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect
