package httpbin

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/gorilla/mux"
)

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

func TestHandleBearer(t *testing.T) {
	token := "some-token"
	target := "http://test.com/basic-auth"
	authHeader := fmt.Sprintf("Bearer %s", token)
	headers := map[string][]string{"Authorization": []string{authHeader}}
	req := newTestRequest(emptyServer.handleBearer(), target, "GET", testReqHeaders(headers))
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	testCases := jsonAssertion{
		{"token", token},
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
	expectedResponseKeys := []string{"authenticated", "token"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}
