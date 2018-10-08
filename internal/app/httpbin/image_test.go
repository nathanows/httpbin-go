package httpbin

import (
	"os"
	"testing"
)

func TestHandleImage(t *testing.T) {
	target := "http://test.com/image"
	headers := map[string][]string{"accept": []string{"image/png"}}
	req := newTestRequest(reqInspectServer.handleImage(""), target, "GET", testReqHeaders(headers))
	os.Chdir("../../..")
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}
	os.Chdir("internal/app/httpbin")

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if len(req.rawResponse) <= 0 {
		t.Error("Response should be non-empty")
	}

	if val := req.response.Header().Get("Content-Type"); val != "image/png" {
		t.Errorf("Expected returned Content-Type to be image/png, got: %v", val)
	}
}

func TestHandleImage_DirectRoute(t *testing.T) {
	target := "http://test.com/image/jpeg"
	headers := map[string][]string{"accept": []string{"image/png"}}
	req := newTestRequest(reqInspectServer.handleImage("image/jpeg"), target, "GET", testReqHeaders(headers))
	os.Chdir("../../..")
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}
	os.Chdir("internal/app/httpbin")

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if len(req.rawResponse) <= 0 {
		t.Error("Response should be non-empty")
	}

	if val := req.response.Header().Get("Content-Type"); val != "image/jpeg" {
		t.Errorf("Expected returned Content-Type to be image/jpeg, got: %v", val)
	}
}
