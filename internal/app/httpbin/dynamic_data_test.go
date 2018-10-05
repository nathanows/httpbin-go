package httpbin

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
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
	code := 304
	target := fmt.Sprintf("http://test.com/drip?duration=%d&numbytes=%d&code=%d&delay=%d", duration, numbytes, code, delay)
	req := newTestRequest(authServer.handleDrip(), target, "GET", testReqStatus([]int{code}))

	start := time.Now()
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}
	elapsedMicroseconds := time.Since(start) / time.Microsecond

	if elapsedMicroseconds < time.Duration(delay+duration) {
		t.Errorf("Request should take at least as long delay + duration...")
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if len(req.rawResponse) != numbytes {
		t.Errorf("expected response to be %d length, got %d", len(req.rawResponse), numbytes)
	}
}

func TestHandleLinks(t *testing.T) {
	links := 2
	offset := 0
	target := fmt.Sprintf("http://test.com/links/%d/%d", links, offset)
	req := newTestRequest(authServer.handleLinks(), target, "GET")
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"n": strconv.Itoa(links), "offset": strconv.Itoa(offset)})

	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if headerVal := req.response.Header().Get("Content-Type"); headerVal != "text/html; charset=utf-8" {
		t.Errorf("Content-Type should be %s, got: %s", "text/html; charset=utf-8", headerVal)
	}

	if !strings.Contains(string(req.rawResponse), "/links/2/1") {
		t.Errorf("Should return correct links")
	}
}

func TestHandleRange(t *testing.T) {
	numbytes := 5
	target := fmt.Sprintf("http://test.com/range/%d", numbytes)
	duration := 5000
	headers := map[string][]string{"Range": []string{"bytes=1-3"}, "duration": []string{strconv.Itoa(duration)}}
	req := newTestRequest(authServer.handleRange(), target, "GET", testReqStatus([]int{206}), testReqHeaders(headers))
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"numbytes": strconv.Itoa(numbytes)})

	start := time.Now()
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}
	elapsedMicroseconds := time.Since(start) / time.Microsecond

	expectedTime := duration * (3 / 5)
	if elapsedMicroseconds < time.Duration(expectedTime) {
		t.Errorf("Request should take at least as long delay + duration...")
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if string(req.rawResponse) != "bcd" {
		t.Errorf("Should return bytes 1-3, got: %v\n", string(req.rawResponse))
	}

	if headerVal := req.response.Header().Get("Content-Length"); headerVal != "3" {
		t.Errorf("Content-Length should be %s, got: %s", "3", headerVal)
	}
}
