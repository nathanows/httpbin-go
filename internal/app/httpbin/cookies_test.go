package httpbin

import (
	"net/http"
	"testing"
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
