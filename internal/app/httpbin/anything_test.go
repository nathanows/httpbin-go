package httpbin

import (
	"testing"
)

var anythingServer = &Server{}

func TestAnything(t *testing.T) {
	target := "http://test.com/anything/test?something=post"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	req := newTestRequest(anythingServer.handleAnything(), target, "POST", testReqHeaders(headers))
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
