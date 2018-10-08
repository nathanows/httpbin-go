package httpbin

import (
	"fmt"
	"testing"

	"github.com/gorilla/mux"
)

func (s *Server) TestHandleRedirectTo(t *testing.T) {
	url := "https://www.google.com"
	target := fmt.Sprintf("http://test.com/redirect-to?url=%s&status_code=306", url)
	req := newTestRequest(reqInspectServer.handleCookiesDelete(), target, "GET", testReqStatus([]int{306}))
	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"url": url, "status_code": "306"})
	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if val := req.response.Header().Get("Location"); val != url {
		t.Errorf("Expected Location to be %s, got: %s", url, val)
	}
}
