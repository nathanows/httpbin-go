package httpbin

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ndwhtlssthr/httpbin-go/pkg/jsonparser"
)

var emptyServer = &Server{}
var possibleResponseFields = []string{"args", "authenticated", "data", "files", "form", "headers", "json", "method", "origin", "url", "user", "user-agent"}

type jsonAssertion []struct {
	jsonPath string
	expected string
}

type jsonBoolAssertion []struct {
	jsonPath string
	expected bool
}

func TestHandleDelete(t *testing.T) {
	target := "http://test.com/delete?something=good"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	req := newTestRequest(emptyServer.handleDelete(), target, "DELETE", testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"args.something", "good"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"args", "data", "files", "form", "headers", "json", "origin", "url"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleGet(t *testing.T) {
	target := "http://test.com/get?something=else"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	req := newTestRequest(emptyServer.handleGet(), target, "GET", testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"args.something", "else"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"args", "headers", "origin", "url"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandlePatch(t *testing.T) {
	target := "http://test.com/patch?something=patched"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	req := newTestRequest(emptyServer.handlePatch(), target, "PATCH", testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"args.something", "patched"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"args", "data", "files", "form", "headers", "json", "origin", "url"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandlePut(t *testing.T) {
	target := "http://test.com/put?something=put"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	req := newTestRequest(emptyServer.handlePut(), target, "PUT", testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"args.something", "put"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"args", "data", "files", "form", "headers", "json", "origin", "url"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandlePost(t *testing.T) {
	target := "http://test.com/post?something=post"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	req := newTestRequest(emptyServer.handlePost(), target, "POST", testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"args.something", "post"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"args", "data", "files", "form", "headers", "json", "origin", "url"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleAnything(t *testing.T) {
	target := "http://test.com/anything/test?something=post"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	req := newTestRequest(emptyServer.handleAnything(), target, "POST", testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"args.something", "post"},
		{"headers.Accept", "*/*"},
		{"url", target},
		{"method", "POST"},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"args", "data", "files", "form", "headers", "json", "method", "origin", "url"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleStatusCodes(t *testing.T) {
	codeOpts := []int{200, 201, 500}
	codeOptsString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(codeOpts)), ","), "[]")
	target := fmt.Sprintf("http://test.com/status/%s", codeOptsString)

	req := newTestRequest(emptyServer.handleStatusCodes(), target, "GET", testReqStatus(codeOpts))

	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"codes": codeOptsString})

	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if req.rawJSON != nil {
		t.Errorf("Response body should be empty, got: %v", string(req.rawJSON))
	}
}

func TestHandleHeaders(t *testing.T) {
	target := "http://test.com/headers"
	headers := map[string][]string{"Accept": []string{"*/*"}, "Something-Else": []string{"one", "two"}}
	req := newTestRequest(emptyServer.handleHeaders(), target, "GET", testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"headers.Accept", "*/*"},
		{"headers.Something-Else", "one,two"},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"headers"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleIP(t *testing.T) {
	target := "http://test.com/ip"
	req := newTestRequest(emptyServer.handleIP(), target, "GET")
	origin := "123.4.5.6"
	req.baseRequest.RemoteAddr = origin
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"origin", origin},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"origin"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleIP_XForwardedFor(t *testing.T) {
	target := "http://test.com/ip"
	forwardedOrigin := "1.1.1.1"
	headers := map[string][]string{"X-Forwarded-For": []string{forwardedOrigin}}
	req := newTestRequest(emptyServer.handleIP(), target, "GET", testReqHeaders(headers))
	req.baseRequest.RemoteAddr = "123.4.5.6"
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"origin", forwardedOrigin},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"origin"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleUserAgent(t *testing.T) {
	target := "http://test.com/user-agent"
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
	headers := map[string][]string{"User-Agent": []string{userAgent}}
	req := newTestRequest(emptyServer.handleUserAgent(), target, "GET", testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"user-agent", userAgent},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
	expectedResponseKeys := []string{"user-agent"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleBasicAuth(t *testing.T) {
	user := "steve"
	pass := "s3cr3t"
	target := fmt.Sprintf("http://test.com/basic-auth/%s/%s", user, pass)
	base64Auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pass)))
	authHeader := fmt.Sprintf("Basic %s", base64Auth)
	headers := map[string][]string{"Authorization": []string{authHeader}}
	req := newTestRequest(emptyServer.handleBasicAuth(), target, "GET", testReqHeaders(headers))
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"user": user, "password": pass})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"user", user},
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}

	if val := req.parsedJSON.Path("authenticated").Data(); val != true {
		t.Errorf("Expected 'authenticated' to be 'true', got: %v", val)
	}
	expectedResponseKeys := []string{"authenticated", "user"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
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
