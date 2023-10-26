package request

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/go-faker/faker/v4"
	"github.com/jinzhu/copier"
)

type call struct {
	Service string
	Method string
	Arguments []interface{}
	Response interface{}
	Run func(*Request)
	onlyOnce bool
}

// When matches calls only when specific arguments are passed to the service and method
func (c *call) When(arguments ...interface{}) *call {
	for _, arg := range arguments {
		if v := reflect.ValueOf(arg); v.Kind() == reflect.Func {
			panic(fmt.Sprintf("cannot use Func in expectations. Use request.AnythingOfType(\"%T\")", arg))
		}
	}
	c.Arguments = arguments
	return c
}
// Return accepts a struct to return as the value to the call
func (c *call) Return(d interface{}) *call {
	c.Response = d
	return c
}

func (c *call) Do(r func(*Request)) *call {
	c.Run = r
	return c
}
func (c *call) Once() *call {
	c.onlyOnce = true
	return c
}
func (c *call) IsOnce() bool {
	return c.onlyOnce
}

type mock struct {
	ExpectedCalls map[string]map[string][]*call
}

var instance *mock
var once sync.Once

// GetMock returns a thread safe mock struct singleton to intercept calls
func GetMock() *mock {
	once.Do(func() {
			instance = &mock{
				ExpectedCalls: make(map[string]map[string][]*call),
			}
	})
	return instance
}

func getType(t interface{}) reflect.Type {
	v := reflect.ValueOf(t)
	if reflect.TypeOf(t).Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Type()
}

func (m *mock) SendHook (r *Request) {
	fmt.Printf("service: %s, operation: %s\n", r.ClientInfo.ServiceName, r.Operation.Name)
	// TODO: Record every operation
  calls := []*call{}
	if _, ok := m.ExpectedCalls["*"]; ok {
		// pattern match *.*
		if c, ok := m.ExpectedCalls["*"]["*"]; ok {
			calls = append(calls, c...)
		}
	}
	if _, ok := m.ExpectedCalls[r.ClientInfo.ServiceName]; ok {
		if c, ok := m.ExpectedCalls[r.ClientInfo.ServiceName][r.Operation.Name]; ok {
			calls = append(calls, c...)
		}
		// pattern match Service.*
		if c, ok := m.ExpectedCalls[r.ClientInfo.ServiceName]["*"]; ok {
			calls = append(calls, c...)
		}
	}

	if err := m.Fake(r.Data); err != nil {
		fmt.Println("Error creating fake data", err)
	}
	for _, call := range calls {
		if call.Response != nil {
			m.Copy(call.Response, r.Data)
			// if call.IsOnce() {
			// 	// TODO: Remove from the expected call list on the mock forever
			// }
		}
		if call.Run !=  nil {
			call.Run(r)
		}
	}
}
// Copy deep copies from one interface to the other. Ignores nil fields and fields that do not exist in target (to). Target must be a pointer
func (m *mock) Copy(from interface{}, to interface{}) {
	copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}

// Fake generates fake data from on an empty struct. example: Fake(&s3.PutObjectOutput{})
func (m *mock) Fake(in interface{}) error {
	outputType := getType(in)
	outputPointer := reflect.New(outputType)   // this type of this variable is reflect.Value.
	outputValue := outputPointer.Elem()        // this type of this variable is reflect.Value.
	outputInterface := outputValue.Interface() // this type of this variable is interface{}
	outPtr := &outputInterface
	// fmt.Printf("Output type is %+v\n", outPtr)
	if err := faker.FakeData(outPtr); err != nil {
		return err
	}
	m.Copy(*outPtr, in)
	return nil
	// return *outPtr, nil

	// if err := faker.FakeData(outPtr); err == nil {
	// 	m.Copy(*outPtr, t)
	// } else {
	// 	fmt.Println("Error creating fake data", err)
	// }
}

func (m *mock) On(serviceName string, methodName string) *call {

	// m.mutex.Lock()
	// defer m.mutex.Unlock()
	// c := newCall(m, methodName, assert.CallerInfo(), arguments...)
	c := newCall(serviceName, methodName)
	if _, ok := m.ExpectedCalls[serviceName]; !ok {
		m.ExpectedCalls[serviceName] = make(map[string][]*call)
	}
	if _, ok := m.ExpectedCalls[serviceName][methodName]; !ok {
		m.ExpectedCalls[serviceName][methodName] = []*call{}
	}
	m.ExpectedCalls[serviceName][methodName] = append(m.ExpectedCalls[serviceName][methodName], c)
	return c
}

func newCall(serviceName string, methodName string) *call {
	return &call {
		Service: serviceName,
		Method: methodName,
	}
}

// AnythingOfTypeArgument contains the type of an argument
// for use when type checking.  Used in Diff and Assert.
//
// Deprecated: this is an implementation detail that must not be used. Use [AnythingOfType] instead.
type AnythingOfTypeArgument = anythingOfTypeArgument

// anythingOfTypeArgument is a string that contains the type of an argument
// for use when type checking.  Used in Diff and Assert.
type anythingOfTypeArgument string

// AnythingOfType returns a special value containing the
// name of the type to check for. The type name will be matched against the type name returned by [reflect.Type.String].
//
// Used in Diff and Assert.
//
// For example:
//
//	Assert(t, AnythingOfType("string"), AnythingOfType("int"))
func AnythingOfType(t string) AnythingOfTypeArgument {
	return anythingOfTypeArgument(t)
}
