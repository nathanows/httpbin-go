package httpbin

import "testing"

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
