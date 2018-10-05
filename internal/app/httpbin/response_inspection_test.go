package httpbin

import (
	"fmt"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandleCache_NoHeaders(t *testing.T) {
	target := "http://test.com/cache"
	req := newTestRequest(reqInspectServer.handleCache(), target, "GET", testReqStatus([]int{200}))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if req.response.Header().Get("etag") == "" {
		t.Errorf("etag header should be set")
	}

	if req.response.Header().Get("last-modified") == "" {
		t.Errorf("last-modified header should be set")
	}

	expectedResponseKeys := []string{"args", "headers", "origin", "url"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleCache_Expired(t *testing.T) {
	target := "http://test.com/cache"
	headers := map[string][]string{"If-Modified-Since": []string{"some-date"}}
	req := newTestRequest(reqInspectServer.handleCache(), target, "GET", testReqStatus([]int{304}), testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if req.response.Header().Get("etag") != "" {
		t.Errorf("etag header should not be set")
	}

	if req.response.Header().Get("last-modified") != "" {
		t.Errorf("last-modified header should not be set")
	}

	expectedResponseKeys := []string{}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleCacheControl(t *testing.T) {
	value := "12"
	target := fmt.Sprintf("http://test.com/cache/%s", value)
	req := newTestRequest(reqInspectServer.handleCacheControl(), target, "GET")
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"value": value})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	expectedCacheControlVal := fmt.Sprintf("public, max-age=%s", value)
	if headerVal := req.response.Header().Get("cache-control"); headerVal != expectedCacheControlVal {
		t.Errorf("cache-control header should be %s, got: %s", expectedCacheControlVal, headerVal)
	}

	expectedResponseKeys := []string{"args", "headers", "origin", "url"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleETag_IfNoneMatch(t *testing.T) {
	etag := "some-tag"
	target := fmt.Sprintf("http://test.com/etag/%s", etag)
	headers := map[string][]string{"If-None-Match": []string{etag, "some-other-val"}}
	req := newTestRequest(reqInspectServer.handleETag(), target, "GET", testReqStatus([]int{304}), testReqHeaders(headers))
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"etag": etag})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if headerVal := req.response.Header().Get("etag"); headerVal != etag {
		t.Errorf("etag header should be %s, got: %s", etag, headerVal)
	}

	expectedResponseKeys := []string{}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleETag_IfMatch(t *testing.T) {
	etag := "some-tag"
	target := fmt.Sprintf("http://test.com/etag/%s", etag)
	headers := map[string][]string{"If-Match": []string{"*"}}
	req := newTestRequest(reqInspectServer.handleETag(), target, "GET", testReqStatus([]int{412}), testReqHeaders(headers))
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"etag": etag})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	expectedResponseKeys := []string{}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleETag_Base(t *testing.T) {
	etag := "some-tag"
	target := fmt.Sprintf("http://test.com/etag/%s", etag)
	req := newTestRequest(reqInspectServer.handleETag(), target, "GET")
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"etag": etag})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	expectedResponseKeys := []string{"args", "headers", "origin", "url"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

func TestHandleResponseHeaders(t *testing.T) {
	target := "http://test.com/response-headers?something=test&Another=good"
	headers := map[string][]string{"Pre-Existing": []string{"here"}}
	req := newTestRequest(reqInspectServer.handleResponseHeaders(), target, "POST", testReqHeaders(headers))
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"something": "test", "Another": "good"})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if val := req.parsedJSON.Path("something").String(); val != "test" {
		t.Errorf("Expected something to have value of test, got: %s", val)
	}
	if val := req.parsedJSON.Path("Another").String(); val != "good" {
		t.Errorf("Expected Another to have value of good, got: %s", val)
	}
	if val := req.parsedJSON.Path("Content-Type").String(); val != "application/json" {
		t.Errorf("Expected Content-Type to be application/json, got: %s", val)
	}
}
