package httpbin

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandleCookies(t *testing.T) {
	target := "http://test.com/cookies"
	req := newTestRequest(reqInspectServer.handleCookies(), target, "GET")
	userID := "1234"
	req.baseRequest.AddCookie(&http.Cookie{Name: "USER_ID", Value: userID})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if val := req.parsedJSON.Path("cookies.USER_ID").String(); val != userID {
		t.Errorf("Expected USER_ID in response to eq: %s, got: %s", userID, val)
	}

	expectedResponseKeys := []string{"cookies"}
	if err := req.validateCorrectFields(expectedResponseKeys); err != nil {
		t.Errorf("Incorrect response keys returned. Failure: %v", err)
	}
}

// TODO - Need to put more work in to tests for handleCookiesDelete and handleCookiesSet
// currently just testing that these requests are redirecting...

func TestHandleCookiesDelete(t *testing.T) {
	target := "http://test.com/cookies/delete?test"
	req := newTestRequest(reqInspectServer.handleCookiesDelete(), target, "GET", testReqStatus([]int{302}))
	req.baseRequest.AddCookie(&http.Cookie{Name: "USER_ID", Value: "1234"})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
}

func TestHandleCookiesSet_QueryParams(t *testing.T) {
	target := "http://test.com/cookies/set?test=val"
	req := newTestRequest(reqInspectServer.handleCookiesSet(), target, "GET", testReqStatus([]int{302}))
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"test": "val"})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
}

func TestHandleCookiesSet_URLParams(t *testing.T) {
	target := "http://test.com/cookies/set/test/val"
	req := newTestRequest(reqInspectServer.handleCookiesSet(), target, "GET", testReqStatus([]int{302}))
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"name": "test", "value": "val"})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}
}
