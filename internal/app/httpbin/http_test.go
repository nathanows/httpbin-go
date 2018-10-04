package httpbin

import "testing"

var httpServer = &Server{}

func TestHandleDelete(t *testing.T) {
	target := "http://test.com/delete?something=good"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	req := newTestRequest(httpServer.handleDelete(), target, "DELETE", testReqHeaders(headers))
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
	req := newTestRequest(httpServer.handleGet(), target, "GET", testReqHeaders(headers))
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
	req := newTestRequest(httpServer.handlePatch(), target, "PATCH", testReqHeaders(headers))
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
	req := newTestRequest(httpServer.handlePut(), target, "PUT", testReqHeaders(headers))
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
	req := newTestRequest(httpServer.handlePost(), target, "POST", testReqHeaders(headers))
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
