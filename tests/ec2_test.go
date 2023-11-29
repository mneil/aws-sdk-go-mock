// File really does nothing. We're patching the aws sdk. This is purely a "test" entrypoint
package tests

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func TestEc2RunInstances(t *testing.T) {
	sess := session.Must(session.NewSession())

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	svc := ec2.New(sess)

	mock := request.GetMock()
	mock.On("ec2", "RunInstances").Return(&struct{}{})

	reservation, err := svc.RunInstances(&ec2.RunInstancesInput{})
	if err != nil {
		t.Fatalf(`TestEc2RunInstances: RunInstances error %v, want nil`, err)
	}

	instance, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: []*string{reservation.Instances[0].ImageId},
	})

	if err != nil {
		t.Fatalf(`TestEc2RunInstances: DescribeInstances error %v, want nil`, err)
	}
	if len(instance.Reservations) == 0 {
		t.Fatalf("TestEc2RunInstances: instance has no reservations")
	}

	if *(instance.Reservations[0].Instances[0].ImageId) != *(reservation.Instances[0].ImageId) {
		t.Fatalf(`TestEc2RunInstances: DescribeInstances have %s, want %s`, *(instance.Reservations[0].Instances[0].ImageId), *(reservation.Instances[0].ImageId))
	}


}
