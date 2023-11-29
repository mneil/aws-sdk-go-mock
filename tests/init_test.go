package tests

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/go-faker/faker/v4/pkg/options"
)

func init() {
	mock :=request.GetMock()
	mock.FakeOptions(
		options.WithRandomMapAndSliceMaxSize(1),
		options.WithRandomMapAndSliceMinSize(1),
	)
}
