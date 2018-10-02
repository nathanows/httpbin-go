package httpbin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ndwhtlssthr/httpbin-go/pkg/jsonparser"
)

func TestHandleDelete(t *testing.T) {
	target := "http://test.com/delete?something=good"
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		t.Errorf("Unable to build request. Err: %v", err)
	}
	req.Header = map[string][]string{"Accept": []string{"*/*"}}

	rr := httptest.NewRecorder()

	server := &Server{}
	http.HandlerFunc(server.HandleDelete()).
		ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	jsonParsed, err := jsonparser.ParseJSON(rr.Body.Bytes())
	if err != nil {
		t.Errorf("Unable to parse returned JSON. Err: %v", err)
	}

	testCases := []struct {
		jsonPath string
		expected string
	}{
		{"args.something", "good"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	for _, tc := range testCases {
		if val := jsonParsed.Path(tc.jsonPath).String(); val != tc.expected {
			t.Errorf("Incorrect val after unmarshal; expected: %v, got: %v", tc.expected, val)
		}
	}
}
