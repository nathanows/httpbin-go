package httpbin

import (
	"os"
	"strings"
	"testing"
)

func TestHandleDeny(t *testing.T) {
	target := "http://test.com/deny"
	req := newTestRequest(reqInspectServer.handleDeny(), target, "GET")
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if headerVal := req.response.Header().Get("Content-Type"); headerVal != "text/plain" {
		t.Errorf("Content-Type should be %s, got: %s", "text/plain", headerVal)
	}

	if !strings.Contains(string(req.rawResponse), "YOU SHOULDN'T BE HERE") {
		t.Errorf("Should return denial message")
	}
}

func TestHandleEncodingUTF8(t *testing.T) {
	os.Chdir("../../..")
	target := "http://test.com/encoding/utf8"
	req := newTestRequest(reqInspectServer.handleEncodingUTF8(), target, "GET")
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}
	os.Chdir("internal/app/httpbin")

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if headerVal := req.response.Header().Get("Content-Type"); headerVal != "text/html; charset=utf-8" {
		t.Errorf("Content-Type should be %s, got: %s", "text/html; charset=utf-8", headerVal)
	}

	if !strings.Contains(string(req.rawResponse), "UTF-8 encoded sample plain-text file") {
		t.Errorf("Should be the UTF-8 sample file")
	}
}

func TestHandleHTML(t *testing.T) {
	os.Chdir("../../..")
	target := "http://test.com/html"
	req := newTestRequest(reqInspectServer.handleHTML(), target, "GET")
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}
	os.Chdir("internal/app/httpbin")

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if headerVal := req.response.Header().Get("Content-Type"); headerVal != "text/html; charset=utf-8" {
		t.Errorf("Content-Type should be %s, got: %s", "text/html; charset=utf-8", headerVal)
	}

	if !strings.Contains(string(req.rawResponse), "Melville") {
		t.Errorf("Should render templates/moby.html")
	}
}

func TestHandleJSON(t *testing.T) {
	target := "http://test.com/json"
	req := newTestRequest(reqInspectServer.handleJSON(), target, "GET")
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	testCases := jsonAssertion{
		{"title", "Sample Slide Show"},
		{"date", "date of publication"},
	}

	if err := req.runTestCases(testCases); err != nil {
		t.Errorf("Failed test case. Failure: %v", err)
	}
}

func TestHandleRobotsTxt(t *testing.T) {
	target := "http://test.com/robots.txt"
	req := newTestRequest(reqInspectServer.handleRobotsTxt(), target, "GET")
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if headerVal := req.response.Header().Get("Content-Type"); headerVal != "text/plain" {
		t.Errorf("Content-Type should be %s, got: %s", "text/plain", headerVal)
	}

	if !strings.Contains(string(req.rawResponse), "Disallow") {
		t.Errorf("Should return robot.txt message")
	}
}

func TestHandleXML(t *testing.T) {
	os.Chdir("../../..")
	target := "http://test.com/encoding/xml"
	req := newTestRequest(reqInspectServer.handleXML(), target, "GET")
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}
	os.Chdir("internal/app/httpbin")

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if headerVal := req.response.Header().Get("Content-Type"); headerVal != "application/xml" {
		t.Errorf("Content-Type should be %s, got: %s", "application/xml", headerVal)
	}

	if !strings.Contains(string(req.rawResponse), "Yours Truly") {
		t.Errorf("Should be the XML sample file")
	}
}
