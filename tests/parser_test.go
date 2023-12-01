// File really does nothing. We're patching the aws sdk. This is purely a "test" entrypoint
package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/parser"
)

type TestParseError struct {
	s string
	err error
}
func (e *TestParseError) Error() string {
	return fmt.Sprintf("String: %s\nError: %s", e.s, e.err.Error())
}

func NewTestParseError(text string, err error) error {
	return &TestParseError{text, err}
}

func parse(search string) (*parser.JqLite, error) {
	queryParser := parser.Parser()
	query, err := queryParser.ParseString("", search)
	if err != nil {
		return nil, NewTestParseError(search, err)
	} else {
		// debug query
		// s, _ := json.MarshalIndent(query, "", "\t")
		// fmt.Printf("Query: %s\n%s\n\n", search, s)
	}
	return query, nil
}


func TestParser(t *testing.T) {
	errors := make([]error, 0)
	cases := []string {
		"instance",
		"instance.InstanceId",
		"instance[]",
		"instance[][0]",
		"instance.InstanceId.Something",
		"instance[].InstanceId",
		"instance[7].InstanceId",
		"instance.InstanceId[1].Something",
		"instance.InstanceId[1].Something",
	}

	for _, testCase := range cases {
		if _, err := parse(testCase); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		out := ""
		for _, testErr := range errors {
			out += fmt.Sprintf("%s\n", testErr.Error())
		}
		t.Fatalf("TestParser:\n%s", out)
	}
}
