package httpbin

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandleStatusCodes(t *testing.T) {
	codeOpts := []int{200, 201, 500}
	codeOptsString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(codeOpts)), ","), "[]")
	target := fmt.Sprintf("http://test.com/status/%s", codeOptsString)

	req := newTestRequest(emptyServer.handleStatusCodes(), target, "GET", testReqStatus(codeOpts))

	req.baseRequest = mux.SetURLVars(req.baseRequest, map[string]string{"codes": codeOptsString})

	if err := req.make(); err != nil {
		t.Errorf("Failed to make request. Err: %v", err)
	}

	if err := req.validateStatusCode(); err != nil {
		t.Errorf("Failed request base validations. Failure: %v", err)
	}

	if req.rawJSON != nil {
		t.Errorf("Response body should be empty, got: %v", string(req.rawJSON))
	}
}
