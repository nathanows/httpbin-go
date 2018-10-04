package httpbin

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/ndwhtlssthr/httpbin-go/pkg/jsonparser"
)

var possibleResponseFields = []string{"args", "authenticated", "data", "files", "form", "headers", "json", "method", "origin", "token", "url", "user", "user-agent"}

type jsonAssertion []struct {
	jsonPath string
	expected string
}

type jsonBoolAssertion []struct {
	jsonPath string
	expected bool
}

type testRequest struct {
	baseRequest *http.Request

	// inputs
	handlerFunc http.HandlerFunc
	target      string
	method      string
	headers     map[string][]string
	status      []int // with multiple status codes a random status is returned

	// output
	rawJSON    []byte
	parsedJSON *jsonparser.Container
	response   *httptest.ResponseRecorder
}

func testReqMethod(method string) func(*testRequest) {
	return func(tr *testRequest) {
		tr.method = method
	}
}

func testReqHeaders(headers map[string][]string) func(*testRequest) {
	return func(tr *testRequest) {
		tr.headers = headers
	}
}

func testReqStatus(codes []int) func(*testRequest) {
	return func(tr *testRequest) {
		tr.status = codes
	}
}

func newTestRequest(handlerFunc http.HandlerFunc, target, method string, opts ...func(*testRequest)) *testRequest {
	tr := &testRequest{
		handlerFunc: handlerFunc,
		target:      target,
		method:      method,
		headers:     make(map[string][]string),
		status:      []int{http.StatusOK},
	}

	for _, opt := range opts {
		opt(tr)
	}

	req, _ := http.NewRequest(tr.method, tr.target, nil)
	tr.baseRequest = req

	return tr
}

func (tr *testRequest) make() error {
	tr.baseRequest.Header = tr.headers

	rr := httptest.NewRecorder()

	tr.handlerFunc.ServeHTTP(rr, tr.baseRequest)

	tr.response = rr

	tr.rawJSON = rr.Body.Bytes()

	if rr.Body.Bytes() != nil {
		jsonParsed, err := jsonparser.ParseJSON(rr.Body.Bytes())
		if err != nil {
			return fmt.Errorf("Unable to parse returned JSON. Err: %v", err)
		}
		tr.parsedJSON = jsonParsed
	}

	return nil
}

func (tr *testRequest) validateStatusCode() error {
	// status in provided statuses
	var validStatus = false
	for _, code := range tr.status {
		if code == tr.response.Code {
			validStatus = true
		}
	}
	if !validStatus {
		return fmt.Errorf("Status code differs. Expected code to be one of %v, got %d", tr.status, tr.response.Code)
	}

	return nil
}

func (tr *testRequest) runTestCases(testCases jsonAssertion) error {
	for _, tc := range testCases {
		if val := tr.parsedJSON.Path(tc.jsonPath).String(); val != tc.expected {
			return fmt.Errorf("Incorrect val after unmarshal; for %v, expected: %v, got: %v", tc.jsonPath, tc.expected, val)
		}
	}
	return nil
}

func (tr *testRequest) validateCorrectFields(expected []string) error {
	for _, field := range expected {
		if val := tr.parsedJSON.Path(field).Data(); val == nil {
			return fmt.Errorf("Expected field %s to be included in response", field)
		}
	}

	expectedNotIncluced := sliceDiff(possibleResponseFields, expected)

	for _, field := range expectedNotIncluced {
		if val := tr.parsedJSON.Path(field).Data(); val != nil {
			return fmt.Errorf("%s should not be included in response, got: %s", field, tr.parsedJSON.Path(field).String())
		}
	}
	return nil
}

func sliceDiff(slice1 []string, slice2 []string) []string {
	var diff []string

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return diff
}
