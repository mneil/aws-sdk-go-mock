// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationcostprofileriface provides an interface to enable mocking the AWS Application Cost Profiler service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package applicationcostprofileriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/applicationcostprofiler"
)

// ApplicationCostProfilerAPI provides an interface to enable mocking the
// applicationcostprofiler.ApplicationCostProfiler service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Application Cost Profiler.
//	func myFunc(svc applicationcostprofileriface.ApplicationCostProfilerAPI) bool {
//	    // Make svc.DeleteReportDefinition request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := applicationcostprofiler.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockApplicationCostProfilerClient struct {
//	    applicationcostprofileriface.ApplicationCostProfilerAPI
//	}
//	func (m *mockApplicationCostProfilerClient) DeleteReportDefinition(input *applicationcostprofiler.DeleteReportDefinitionInput) (*applicationcostprofiler.DeleteReportDefinitionOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockApplicationCostProfilerClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ApplicationCostProfilerAPI interface {
	DeleteReportDefinition(*applicationcostprofiler.DeleteReportDefinitionInput) (*applicationcostprofiler.DeleteReportDefinitionOutput, error)
	DeleteReportDefinitionWithContext(aws.Context, *applicationcostprofiler.DeleteReportDefinitionInput, ...request.Option) (*applicationcostprofiler.DeleteReportDefinitionOutput, error)
	DeleteReportDefinitionRequest(*applicationcostprofiler.DeleteReportDefinitionInput) (*request.Request, *applicationcostprofiler.DeleteReportDefinitionOutput)

	GetReportDefinition(*applicationcostprofiler.GetReportDefinitionInput) (*applicationcostprofiler.GetReportDefinitionOutput, error)
	GetReportDefinitionWithContext(aws.Context, *applicationcostprofiler.GetReportDefinitionInput, ...request.Option) (*applicationcostprofiler.GetReportDefinitionOutput, error)
	GetReportDefinitionRequest(*applicationcostprofiler.GetReportDefinitionInput) (*request.Request, *applicationcostprofiler.GetReportDefinitionOutput)

	ImportApplicationUsage(*applicationcostprofiler.ImportApplicationUsageInput) (*applicationcostprofiler.ImportApplicationUsageOutput, error)
	ImportApplicationUsageWithContext(aws.Context, *applicationcostprofiler.ImportApplicationUsageInput, ...request.Option) (*applicationcostprofiler.ImportApplicationUsageOutput, error)
	ImportApplicationUsageRequest(*applicationcostprofiler.ImportApplicationUsageInput) (*request.Request, *applicationcostprofiler.ImportApplicationUsageOutput)

	ListReportDefinitions(*applicationcostprofiler.ListReportDefinitionsInput) (*applicationcostprofiler.ListReportDefinitionsOutput, error)
	ListReportDefinitionsWithContext(aws.Context, *applicationcostprofiler.ListReportDefinitionsInput, ...request.Option) (*applicationcostprofiler.ListReportDefinitionsOutput, error)
	ListReportDefinitionsRequest(*applicationcostprofiler.ListReportDefinitionsInput) (*request.Request, *applicationcostprofiler.ListReportDefinitionsOutput)

	ListReportDefinitionsPages(*applicationcostprofiler.ListReportDefinitionsInput, func(*applicationcostprofiler.ListReportDefinitionsOutput, bool) bool) error
	ListReportDefinitionsPagesWithContext(aws.Context, *applicationcostprofiler.ListReportDefinitionsInput, func(*applicationcostprofiler.ListReportDefinitionsOutput, bool) bool, ...request.Option) error

	PutReportDefinition(*applicationcostprofiler.PutReportDefinitionInput) (*applicationcostprofiler.PutReportDefinitionOutput, error)
	PutReportDefinitionWithContext(aws.Context, *applicationcostprofiler.PutReportDefinitionInput, ...request.Option) (*applicationcostprofiler.PutReportDefinitionOutput, error)
	PutReportDefinitionRequest(*applicationcostprofiler.PutReportDefinitionInput) (*request.Request, *applicationcostprofiler.PutReportDefinitionOutput)

	UpdateReportDefinition(*applicationcostprofiler.UpdateReportDefinitionInput) (*applicationcostprofiler.UpdateReportDefinitionOutput, error)
	UpdateReportDefinitionWithContext(aws.Context, *applicationcostprofiler.UpdateReportDefinitionInput, ...request.Option) (*applicationcostprofiler.UpdateReportDefinitionOutput, error)
	UpdateReportDefinitionRequest(*applicationcostprofiler.UpdateReportDefinitionInput) (*request.Request, *applicationcostprofiler.UpdateReportDefinitionOutput)
}

var _ ApplicationCostProfilerAPI = (*applicationcostprofiler.ApplicationCostProfiler)(nil)
