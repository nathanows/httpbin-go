package httpbin

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestHandleBase64Encode(t *testing.T) {
	value := "some sample value!"
	base64Val := base64.StdEncoding.EncodeToString([]byte(value))
	target := fmt.Sprintf("http://test.com/base64/%s", base64Val)
	req := newTestRequest(authServer.handleBase64Decode(), target, "GET")
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"value": base64Val})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if decoded := string(req.rawResponse); decoded != value {
		t.Errorf("Expected: %s, Got: %s", value, decoded)
	}
}

func TestHandleBytes(t *testing.T) {
	bytes := 7
	target := fmt.Sprintf("http://test.com/bytes/%d", bytes)
	req := newTestRequest(authServer.handleBytes(), target, "GET")
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"n": strconv.Itoa(bytes)})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if respLength := len(req.rawResponse); respLength != bytes {
		t.Errorf("Expected response to be %d random bytes, got %d bytes", bytes, respLength)
	}
}

func TestHandleDelay(t *testing.T) {
	delay := 1000 // using microsecond shortcut for testing
	target := fmt.Sprintf("http://test.com/delay/%d", delay)
	req := newTestRequest(authServer.handleDelay(), target, "GET")
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"delay": strconv.Itoa(delay)})

	start := time.Now()
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}
	elapsedMicroseconds := time.Since(start) / time.Microsecond

	if elapsedMicroseconds < time.Duration(delay) {
		t.Errorf("Request should take at least as long delay...")
	}

	expectedResponseKeys := []string{"url", "args", "form", "data", "origin", "headers", "files"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleDrip(t *testing.T) {
	delay := 1000 // using microsecond shortcut for testing
	duration := 5000
	numbytes := 5
	code := "304"
	target := fmt.Sprintf("http://test.com/drip?duration=%d&numbytes=%d&code=%s&delay=%d", duration, numbytes, code, delay)
	req := newTestRequest(authServer.handleDrip(), target, "GET")

	start := time.Now()
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}
	elapsedMicroseconds := time.Since(start) / time.Microsecond

	if elapsedMicroseconds < time.Duration(delay+duration) {
		t.Errorf("Request should take at least as long delay + duration...")
	}

	if len(req.rawResponse) != numbytes {
		t.Errorf("expected response to be %d length, got %d", len(req.rawResponse), numbytes)
	}
}
