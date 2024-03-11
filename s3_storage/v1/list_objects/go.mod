module example.com/list_objects

go 1.21.1

require github.com/aws/aws-sdk-go v1.46.6

require github.com/jmespath/go-jmespath v0.4.0 // indirect

replace example.com/s3utils => ../s3utils

require example.com/s3utils v1.0.0
