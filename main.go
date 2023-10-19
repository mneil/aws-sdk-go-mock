package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	var bucket, key string
	var timeout time.Duration

	flag.StringVar(&bucket, "b", "", "Bucket name.")
	flag.StringVar(&key, "k", "", "Object key name.")
	flag.DurationVar(&timeout, "d", 0, "Upload timeout.")
	flag.Parse()
	if err := run(bucket, key, timeout); err != nil{
		os.Exit(1)
	}
}
type Rando struct {
	BucketKeyEnabled *bool
	SSEKMSKeyId *string
	SomethingWeird *string
}

func run(bucket string, key string, timeout time.Duration) error {

	// All clients require a Session. The Session provides the client with
 	// shared configuration such as region, endpoint, and credentials. A
 	// Session should be shared where possible to take advantage of
 	// configuration and credential caching. See the session package for
 	// more information.
	sess := session.Must(session.NewSession())

 	// Create a new instance of the service's client with a Session.
 	// Optional aws.Config values can also be provided as variadic arguments
 	// to the New function. This option allows you to provide service
 	// specific configuration.
	svc := s3.New(sess)

	mock := request.GetMock()
	mock.On("s3", "PutObject").Return(&Rando{BucketKeyEnabled: aws.Bool(true), SSEKMSKeyId: aws.String("ABC")})
	mock.On("s3", "PutObject").Return(&Rando{BucketKeyEnabled: aws.Bool(false), SomethingWeird: aws.String("Value")})
	// mock.On("*", "*").Return(&s3.PutObjectOutput{})


	// svc.Handlers.Send.PushFront(func (r *request.Request) {
	// 	fmt.Printf("%s\n", r.Operation.Name)
	// 	// r.Operation.Name == "PutObject"
	// 	// r.Params.Bucket == &"unknown"
	// 	// r.Data (fill response data)
	// })
	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	if cancelFn != nil {
  		defer cancelFn()
	}

	// Uploads the object to S3. The Context will interrupt the request if the
	// timeout expires.
	r, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   os.Stdin,
	})
	fmt.Printf("Returned BucketKeyEnabled:%t, SSEKMSKeyId: %s\n", *r.BucketKeyEnabled, *r.SSEKMSKeyId)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			return fmt.Errorf("upload canceled due to timeout, %v", err)
		} else {
			return fmt.Errorf("failed to upload object, %v", err)
		}
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)
	return nil
}
