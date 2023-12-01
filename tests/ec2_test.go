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
	"github.com/aws/aws-sdk-go/parser"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Out is the data from the input struct that you want to map to the output struct
type Out struct {
	// Key to move to the output. You can use "" to specify all values of the struct
	// Optional - Defaults to all keys
	Key string `json:"key"`
	// Property to place the key(s) into of the output struct
	// Optional - Defaults to the root ""
	Property string `json:"property"`
}
// Store the output into a data "store"
type Store struct {
	// Key to use for the store id (like a primary key). Value must resolve to string or int. Example: "Instances[].InstanceId"
	// Required
	Key string `json:"key"`
	// What property of the output to store. Leave empty to store the entire ouput struct
	// Optional - Defaults to "" or all properties
	Property string `json:"property"`
	// Service the store belongs to
	// Optional - Defaults to the current / same service
	Service string `json:"service"`
	// Store to put the data into. Can be any string. Useful only for mapping to subsequent calls
	// Required.
	Store string `json:"store"`
}
// Use some data from the store in the output value
type Use struct {
	// Key from which property to look for the Key of the store
	// Required - Where to get the key ID from
	Key string `json:"key"`
	// Where to put the data from the store
	// Optional - Defaults to "" which is to apply to the root of the output struct
	Out string `json:"out"`
	// Service the store belongs to
	// Optional - Defaults to current service
	Service string `json:"service"`
	// Store to retrieve data from
	// Required
	Store string `json:"store"`
}

type MockData struct {
	Out []Out `json:"out"`
	Stores []Store `json:"stores"`
	Uses []Use `json:"uses"`
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

func mockCopyOut(r *request.Request, outOpts []Out) {
	mock := request.GetMock()
	for _, out := range outOpts {
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
}

func mockAddStore(r *request.Request, stores []Store) {
	queryParser := parser.Parser()
	for _, store := range stores {
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

		keyQuery, err := queryParser.ParseString("", store.Key)
		if err != nil {
			fmt.Println(err)
		}
		keys := GetValues(r.Data, keyQuery)

		var value interface{}
		if store.Property != "" {
			valueQuery, err := queryParser.ParseString("", store.Property)
			if err != nil {
				fmt.Println(err)
			}
			values := GetValues(r.Data, valueQuery)
			if len(values) > 1 {
				valueQueryJson, _ := json.MarshalIndent(valueQuery, "", "\t")
				panic(fmt.Sprintf("No support for multi value storage. Update your storage property\nservice: %s\noperation: %s\nProperty: %s\nParsed Query:%s", service, r.Operation.Name, store.Property, valueQueryJson))
			}
			fmt.Println("Value to convert to intreface", values[0].Kind())
			value = ValueToInterface(values[0])
		} else {
			value = r.Data
		}
		for _, key := range keys {
			mockStore[key.Elem().String()] = value
		}
	}
}

// {
// 	"Index": null
// }
func GetIndexExpr(x parser.IndexExpr, data reflect.Value) []reflect.Value {
	if x.Index == nil {
		values := make([]reflect.Value, 0)
		for i := 0; i < data.Len(); i++ {
			values = append(values, data.Index(i))
		}
		return values
	}else {
		return []reflect.Value{data.Index(*x.Index)}
	}
}
// {
// 	"Sel": "Instances",
// 	"X": [
// 		{
// 			"Index": null
// 		}
// 	]
// }
func GetSelExpr(expr parser.SelectorExpr, data reflect.Value) []reflect.Value {
	value := data.Elem().FieldByName(expr.Sel)
	// fmt.Printf("Value %+v\n", value)
	if expr.X == nil {
		return []reflect.Value{value}
	}
	values := make([]reflect.Value, 0)
	for _, x := range expr.X {
		v := GetIndexExpr(x, value)
		values = append(values, v...)
	}
	return values
}
// {
// 	"Expr": [
// 		{
// 			"Sel": "Instances",
// 			"X": [
// 				{
// 					"Index": null
// 				}
// 			]
// 		},
// 		{
// 			"Sel": "InstanceId",
// 			"X": null
// 		}
// 	]
// }
func GetValues(data interface{}, query *parser.JqLite) []reflect.Value {
	values := []reflect.Value{reflect.ValueOf(data)}
	for _, expr := range query.Expr {
		for _, value := range values {
			v := GetSelExpr(expr, value)
			for i, vv := range v {
				if i > len(values) {
					values = append(values, vv)
				}else {
					values[i] = vv
				}
			}
		}
	}
	return values
}

func ValueToInterface(value reflect.Value) interface{} {
	var out interface{}
	if value.Kind() == reflect.Struct {
		out = value.Interface()
	} else if value.Kind() == reflect.Pointer {
		out = value.Elem().Interface()
	} else if value.Kind() == reflect.Slice {
		tmpOut := make([]interface{}, 0)
		for i := 0; i < value.Len(); i++ {
			fmt.Println("Value in slice", value.Index(i).Kind())
			tmpOut = append(tmpOut, ValueToInterface(value.Index(i)))
		}
		out = tmpOut
	}
	return out
}

// func (c *EC2) RunInstancesRequest(input *RunInstancesInput) (req *request.Request, output *Reservation) {
// type RunInstancesInput struct
// type Reservation struct
// func (c *EC2) DescribeInstancesRequest(input *DescribeInstancesInput) (req *request.Request, output *DescribeInstancesOutput) {
// type DescribeInstancesOutput struct


func TestEc2RunInstances(t *testing.T) {
	Config := &MockServiceConfig{}
	json.Unmarshal(request.MockServiceJson, Config)
	sess := session.Must(session.NewSession())
	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	client := ec2.New(sess)

	mock := request.GetMock()
	mock.On("*", "*").Do(func (r *request.Request) {
		// Look in config to see if service has mappings
		if svc, ok := (*Config)[r.ClientInfo.ServiceName]; ok {
			// Look at service mappings to see if this operation has any mappings
			if op, ok := svc[r.Operation.Name]; ok {
				// Copy data from op.Out into r.Data as defined in the mapping
				mockCopyOut(r, op.Out)
				mockAddStore(r, op.Stores)


				for _, uses := range op.Uses {

					fmt.Printf("op.Uses %+v\n", uses)

					// Start getting data out of store
					var service string
					if uses.Service != "" {
						service = uses.Service
					}else {
						service = r.ClientInfo.ServiceName
					}
					mockService, ok := MockStorage[service]
					if !ok {
						// Do we 404 here?
						fmt.Println("No service found")
						return
					}
					mockStore, ok := mockService[uses.Store]
					if !ok {
						// Do we 404 here?
						fmt.Println("No store found")
						return
					}
					// find the values out of the store
					queryParser := parser.Parser()
					valueQuery, err := queryParser.ParseString("", uses.Key)
					if err != nil {
						fmt.Println(err)
					}
					values := GetValues(r.Params, valueQuery)
					for _, value := range values {
						sValue, ok := ValueToInterface(value).(string)
						if !ok {
							// TODO: 404?
							continue
						}
						// store value. Now we can insert it into the output somewhere
						storeValue, ok := mockStore[sValue]
						// fmt.Printf("FOUND STORE VALUE TO USE %t, %+v\n", ok, storeValue)
						if !ok {
							// TODO: 404?
							continue
						}

						// mock := request.GetMock()
						// for _, out := range outOpts {
						p := reflect.ValueOf(r.Data)
						f := p.Elem().FieldByName(uses.Out)
						src := storeValue
						// fmt.Println("Found store Value", storeValue)
						// fmt.Println("Writing value out to ", f.Interface())

						fmt.Println("")
						fmt.Println("")

						// TODO: Need to copy data into a slice and figure out how to handle different lengths
						if f.Kind() == reflect.Slice {
							if reflect.TypeOf(src).Kind() != reflect.Slice {
								tmpSlice := make([]interface{}, 1)
								tmpSlice[0] = storeValue
								src = tmpSlice

								// panic(fmt.Sprintf("Trying to replace a slice with a non-slice.\n\tService %s\n\tOperation: %s\n", r.ClientInfo.ServiceName, r.Operation.Name))
							}
							srcX := reflect.New(reflect.TypeOf(src))
							srcX.Elem().Set(reflect.ValueOf(src))

							destX := reflect.New(reflect.TypeOf(f))
							destX.Elem().Set(reflect.ValueOf(f))

							logSrc := srcX.Interface()
							logF := reflect.ValueOf(f).Pointer()

							// fmt.Println("src", logSrc)
							fmt.Println("f", logF)

							fmt.Println("canAddr src", canAddr(logSrc))
							fmt.Println("canAddr f", canAddr(logF))

							// fmt.Println("srcX", srcX)
							// fmt.Println("destX", f.Interface())

							mock.Copy(logSrc, logF)
							// sliceSrc := src.([]interface{})
							// for i := 0; i < len(sliceSrc); i++ {
							// 	// if i > f.Len() {
							// 	// 	fAsInterface := f.Interface().([]interface{})
							// 	// 	f = append(fAsInterface, interface{})
							// 	// }
							// 	mock.Copy(sliceSrc[i], f.Index(i).Interface())
							// }
						}

						// data := r.Data.(*ec2.DescribeInstancesOutput)
						// fmt.Println("Kind of src", canAddr(src))
						// fmt.Println("Kind of f", canAddr(f))
						// fmt.Println("Kind of r.Data.Reservations", canAddr(data))
						fmt.Println("")
						fmt.Println("")

						// fmt.Println("Kind of f", f.Kind())
						// fmt.Println("Kind of src", reflect.ValueOf(src).Kind())
						// fmt.Println("SRC", src)
						// fmt.Println(reflect.ValueOf(&src).Addr())
						mock.Copy(src, f.Interface())
						// mock.Copy(src, f.PointerTo())
						// mock.Copy(src, f.Pointer())
						// mock.Copy(src, f.Interface())
						// mock.Copy(src, f.Addr())
						// mock.Copy(src, f.Elem())
						// }
						// TODO: support copying individual props


					}
					// fmt.Println("op.Uses", uses.Store)

				}


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

	reservation, err := client.RunInstances(&ec2.RunInstancesInput{ImageId: aws.String("abc-123")})

	fmt.Println("Image ID", *reservation.Instances[0].ImageId)

	if err != nil {
		t.Fatalf(`TestEc2RunInstances: RunInstances error %v, want nil`, err)
	}

	instance, err := client.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: []*string{reservation.Instances[0].InstanceId},
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

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

func canAddr(i interface{}) bool {
	return indirect(reflect.ValueOf(i)).CanAddr()
}
