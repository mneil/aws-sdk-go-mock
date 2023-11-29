// File really does nothing. We're patching the aws sdk. This is purely a "test" entrypoint
package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Out is the data from the input struct that you want to map to the output struct
type Out struct {
	// Which key to move to the output. You can use "*" to specify all values of the struct
	Key string `json:"key"`
	// Property to place the key(s) into of the output struct
	Property string `json:"property"`
}
// Store the output into a data "store"
type Store struct {
	// the name of the store to put the data into
	Store string `json:"store"`
	// the property of the output struct to use as the key in the store (like a primary key). Example: "Instances[].InstanceId"
	Key string `json:"key"`
	// What property of the output to store. Leave empty to store the entire ouput struct
	Property string `json:"property"`
}
// Use some data from the store in the output value
type Use struct {
	// Store to retrieve data from
	Store string `json:"key"`
	// From which property to look for the Key of the store
	From string `json:"from"`
}

type MockData struct {
	Out []Out
	Stores []Store
	Uses []Use
}
type MockOperationConfig map[string]MockData
type MockServiceConfig map[string]MockOperationConfig

// {
// 	// Service
// 	"ec2": {
// 		// Input method
// 		"RunInstancesInput": {
// 			// Out is what it returns where. By default we take all keys and lay it over output
// 			"Out": [
// 				{ // return all keys from input struct into the Instances property
// 					"Key": "*",
// 					"Property": "Instances" // ,
// 					// "Type": "list" // struct
// 				}
// 			],
// 			// Store data into a map
// 			"Stores": [
// 				{
// 					// the store to save this data in
// 					"Store": "Reservations"
// 					// what to use as the key in the map
// 					"Key": "Instances[].InstanceId",
// 					// name of the key to store the data under
// 					"Property": "Reservations",

// 				}
// 			]
// 		},
// 		"DescribeInstancesInput": {
// 			"Uses": [
// 				{
// 					"Key": "Reservations",
// 					"From": "InstanceIds",
// 				}
// 			]
// 		}
// 	}
// }


func TestEc2RunInstances(t *testing.T) {
	mockServiceConfig := &MockServiceConfig{}
	json.Unmarshal(request.MockServiceJson, mockServiceConfig)
	fmt.Println("mockServiceConfig", mockServiceConfig)
	sess := session.Must(session.NewSession())

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	svc := ec2.New(sess)

	mock := request.GetMock()
	// mock.On("ec2", "RunInstances").Return(&struct{}{})
	mock.On("*", "*").Do(func (r *request.Request) {
		// out, err := json.Marshal(r.Data.(*aws_ec2.Reservation))
		// if err != nil {
		// 	fmt.Println("Error marshalling: ", err)
		// }
		in := request.GetInterfaceValue(r.Params)
		out := request.GetInterfaceValue(r.Data)

		fmt.Println("Call: ", r.ClientInfo.ServiceName, r.Operation.Name)
		fmt.Println("PARAMS", in.Type().Name(), r.Params)
		fmt.Println("RESPONSE", out.Type().Name(), r.Data)
		// fmt.Println("RESPONSE", string(out[:]))
		fmt.Println("")
		// call my jsfun with args
	})




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
