// File really does nothing. We're patching the aws sdk. This is purely a "test" entrypoint
package tests

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
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
	// Service the store belongs to
	Service string `json:"service"`
	// the name of the store to put the data into
	Store string `json:"store"`
	// the property of the output struct to use as the key in the store (like a primary key). Example: "Instances[].InstanceId"
	Key string `json:"key"`
	// What property of the output to store. Leave empty to store the entire ouput struct
	Property string `json:"property"`
}
// Use some data from the store in the output value
type Use struct {
	// Service the store belongs to
	Service string `json:"service"`
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


// <image-id>
type MockStore map[string]interface{}
// RunInstances
type MockStoreService map[string]MockStore

var MockStorage = make(map[string]MockStoreService)
// ec2
// type MockStorage struct {

// 	store map[string]MockStoreService
// }

// func (ms *MockStorage) Service(s string) {

// }

// func NewMockStorage() *MockStorage {

// 	mss := make(map[string]MockStoreService)
// 	return &MockStorage{
// 		store: mss,
// 	}
// }

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
	Config := &MockServiceConfig{}
	json.Unmarshal(request.MockServiceJson, Config)
	sess := session.Must(session.NewSession())

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	svc := ec2.New(sess)

	mock := request.GetMock()
	mock.On("*", "*").Do(func (r *request.Request) {
		// Look in config to see if service has mappings
		if svc, ok := (*Config)[r.ClientInfo.ServiceName]; ok {
			// Look at service mappings to see if this operation has any mappings
			if op, ok := svc[r.Operation.Name]; ok {
				// Copy data from op.Out into r.Data as defined in the mapping
				for _, out := range op.Out {
					p := reflect.ValueOf(r.Data)
					f := p.Elem().FieldByName(out.Property)
					src := r.Params
					// if out.Key != "" && out.Key != "*" {
					// 	// TODO: If this is true we need to get a specific value out of Params
					// }
					// value we map to is a slice. We map to all entries in the slice
					if f.Kind() == reflect.Slice {
						for i := 0; i < f.Len(); i++ {
							e := f.Index(i)
							mock.Copy(src, e.Interface())
						}
					// value is a struct
					} else if f.Kind() == reflect.Struct {
						mock.Copy(src, f.Interface())
					} else {
						fmt.Println("Unable to map out to value of type", f.Kind())
					}
				}
				for _, store := range op.Stores {
					fmt.Println("op.Stores", store.Key)
					service := store.Service
					if service == "" {
						service = r.ClientInfo.ServiceName
					}
					mockService, ok := MockStorage[service]
					if !ok {
						MockStorage[service] = make(map[string]MockStore)
						mockService = MockStorage[service]
					}
					mockStore, ok := mockService[store.Store]
					if !ok {
						mockService[store.Store] = make(map[string]interface{})
						mockStore = mockService[store.Store]
					}

					mockStore[store.Property] = r.Data

					// fmt.Println("op.Stores", store.Service)
					// fmt.Println("op.Stores", store.Property)
					// fmt.Println("op.Stores", store.Store)
				}
				// for _, uses := range op.Uses {
				// 	fmt.Println("op.Uses", uses.From)
				// 	// fmt.Println("op.Uses", uses.Store)
				// }


			}
		}

		// out, err := json.Marshal(r.Data.(*aws_ec2.Reservation))
		// if err != nil {
		// 	fmt.Println("Error marshalling: ", err)
		// }
		// in := request.GetInterfaceValue(r.Params)
		// out := request.GetInterfaceValue(r.Data)

		// fmt.Println("Call: ", r.ClientInfo.ServiceName, r.Operation.Name)
		// fmt.Println("PARAMS", in.Type().Name(), r.Params)
		// fmt.Println("RESPONSE", out.Type().Name(), r.Data)
		// fmt.Println("RESPONSE", string(out[:]))
		fmt.Println("")
		// call my jsfun with args
	})



	// svc.RunInstancesWithContext(context.Background(), &ec2.RunInstancesInput{ImageId: aws.String("abc-123")})

	reservation, err := svc.RunInstances(&ec2.RunInstancesInput{ImageId: aws.String("abc-123")})

	fmt.Println("Image ID", *reservation.Instances[0].ImageId)

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
