module github.com/mneil/aws-sdk-go-mock

go 1.20

replace github.com/aws/aws-sdk-go => ./aws/aws-sdk-go

require (
	github.com/aws/aws-sdk-go v1.45.25
	github.com/jinzhu/copier v0.4.0
)

require (
	github.com/go-faker/faker/v4 v4.2.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
