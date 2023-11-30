package request

import (
	_ "embed"
	"fmt"
	"reflect"
	"sync"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/jinzhu/copier"
)

//go:embed mock_service.json
var MockServiceJson []byte

type call struct {
	Service string
	Method string
	Arguments []interface{}
	Response interface{}
	Run func(*Request)
	onlyOnce bool
	fakeOptions []options.OptionFunc
}

func (c *call) WithFakeOptions(o ...options.OptionFunc) *call {
	c.fakeOptions = o
	return c
}

func (c *call) FakeOptions() []options.OptionFunc {
	return c.fakeOptions
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

type RecordedCall struct{
	Service string
	Operation string
	Params interface{}
	Response interface{}
}

type mock struct {
	ExpectedCalls map[string]map[string][]*call
	calls []RecordedCall
	fakeOptions []options.OptionFunc
}

var instance *mock
var once sync.Once

// GetMock returns a thread safe mock struct singleton to intercept calls
func GetMock() *mock {
	once.Do(func() {
			instance = &mock{
				ExpectedCalls: make(map[string]map[string][]*call),
				calls: make([]RecordedCall, 0),
				fakeOptions: make([]options.OptionFunc, 0),
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

func (m *mock) FakeOptions(o ...options.OptionFunc) {
	m.fakeOptions = o
}

func (m *mock) SendHook (r *Request) {
	fmt.Printf("service: %s, operation: %s\n", r.ClientInfo.ServiceName, r.Operation.Name)
	// TODO: Record every operation
	// find all calls that the user has mocked
  calls := []*call{}
	if _, ok := m.ExpectedCalls["*"]; ok {
		// pattern match *.*
		if c, ok := m.ExpectedCalls["*"]["*"]; ok {
			calls = append(calls, c...)
		}
	}
	// service scoped mocks
	if _, ok := m.ExpectedCalls[r.ClientInfo.ServiceName]; ok {
		// pattern match Service.Opereration
		if c, ok := m.ExpectedCalls[r.ClientInfo.ServiceName][r.Operation.Name]; ok {
			calls = append(calls, c...)
		}
		// pattern match Service.*
		if c, ok := m.ExpectedCalls[r.ClientInfo.ServiceName]["*"]; ok {
			calls = append(calls, c...)
		}
	}
	// get all options set on individual matching calls
	// options set later take precedence
	overrideOptions := make([]options.OptionFunc, 0)
	for _, call := range calls {
		if len(call.fakeOptions) > 0 {
			overrideOptions = append(overrideOptions, call.fakeOptions...)
		}
	}

	// fill the output (r.Data) with fake data based on the shape of r.Data
	if err := m.Fake(r.Data, overrideOptions...); err != nil {
		fmt.Println("Error creating fake data", err)
	}
	// overwrite anything in r.Data with the information returned by user mocks
	// higher fidelity mocks take precedent (ie: *.* is overwritten by Service.* data)
	for _, call := range calls {
		// if len(call.fakeOptions) > 0 {
		// 	m.Fake(r.Data, call.fakeOptions...)
		// }
		if call.Response != nil {
			m.Copy(call.Response, r.Data)
		}
		if call.Run !=  nil {
			call.Run(r)
		}
		// if call.IsOnce() {
		// 	// TODO: Remove from the expected call list on the mock forever
		// }
	}
	m.called(r)
}

// Record a call to the mock
func (m *mock) called(r *Request) {
	m.calls = append(m.calls, RecordedCall{
		Service: r.ClientInfo.ServiceName,
		Operation: r.Operation.Name,
		Params: r.Params,
		Response: r.Data,
	})
}
// Calls returns calls recorded by the SDK
func (m *mock) Calls() []RecordedCall {
	return m.calls
}
// ResetCalls wipes all recorded calls resetting it back to an empty slice
func (m *mock) ResetCalls() {
	m.calls = make([]RecordedCall, 0)
}
// Reset resets the whole mock. calls and ExpectedCalls are removed
func (m *mock) Reset() {
	m.ResetCalls()
	m.ExpectedCalls = make(map[string]map[string][]*call)
}

// Copy deep copies from one interface to the other. Ignores nil fields and fields that do not exist in target (to). Target must be a pointer
func (m *mock) Copy(from interface{}, to interface{}) {
	copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}

func GetInterfaceValue(in interface{}) reflect.Value {
	outputType := getType(in) // get the type of interface regardless if interface is value or pointer
	outputPointer := reflect.New(outputType)   // this type of this variable is reflect.Value.
	return outputPointer.Elem() // gets the value that the output pointer points to
}

// Fake generates fake data from on an empty struct. example: Fake(&s3.PutObjectOutput{})
func (m *mock) Fake(in interface{}, callOptions ...options.OptionFunc) error {
	outputValue := GetInterfaceValue(in)        // this type of this variable is reflect.Value.
	outputInterface := outputValue.Interface() // this type of this variable is interface{}
	outPtr := &outputInterface
	// fmt.Printf("Output type is %+v\n", outPtr)
	opts := make([]options.OptionFunc, 0)
	opts = append(opts, m.fakeOptions...)
	opts = append(opts, callOptions...)
	if err := faker.FakeData(outPtr, opts...); err != nil {
		return err
	}
	m.Copy(*outPtr, in)
	return nil
}

func (m *mock) On(serviceName string, methodName string) *call {
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
		fakeOptions: make([]options.OptionFunc, 0),
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
