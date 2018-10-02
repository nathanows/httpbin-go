package httpbin

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ndwhtlssthr/httpbin-go/pkg/jsonparser"
)

func TestParseRequest_URL_AbsTarget(t *testing.T) {
	target := "http://hbg.com/delete?some_param=2"
	r := httptest.NewRequest("DELETE", target, nil)
	hbr, err := parseRequest(r)
	if err != nil {
		t.Errorf("Failed to ParseRequest. Err: %v", err)
	}

	if hbr.URL != target {
		t.Errorf("got %s, want %s", hbr.URL, target)
	}
}

func TestParseRequest_URL_RelTarget(t *testing.T) {
	target := "http://localhost:8080/delete?test=test2"
	url := &url.URL{Path: "/delete", RawQuery: "test=test2"}
	r := &http.Request{
		URL:  url,
		Host: "localhost:8080",
	}
	hbr, err := parseRequest(r)
	if err != nil {
		t.Errorf("Failed to ParseRequest. Err: %v", err)
	}

	if hbr.URL != target {
		t.Errorf("got %s, want %s", hbr.URL, target)
	}
}

func TestParseRequest_Headers(t *testing.T) {
	target := "http://hbg.com/get"
	r := httptest.NewRequest("GET", target, nil)

	headers := make(map[string][]string)
	headers["test"] = []string{"test1", "test2"}
	headers["Accept"] = []string{"*/*"}
	r.Header = headers

	hbr, err := parseRequest(r)
	if err != nil {
		t.Errorf("Failed to ParseRequest. Err: %v", err)
	}

	if hbr.Headers["test"] != "test1,test2" {
		t.Errorf("got %s, want %s", hbr.Headers["test"], "test1,test2")
	}
	if hbr.Headers["Accept"] != "*/*" {
		t.Errorf("got %s, want %s", hbr.Headers["Accept"], "*/*")
	}
}

func TestParseRequest_Origin(t *testing.T) {
	target := "http://hbg.com/get"
	r := httptest.NewRequest("GET", target, nil)

	hbr, err := parseRequest(r)
	if err != nil {
		t.Errorf("Failed to ParseRequest. Err: %v", err)
	}

	if hbr.Origin != r.RemoteAddr {
		t.Errorf("got %s, want %s", hbr.Origin, r.RemoteAddr)
	}
}

func TestParseRequest_Origin_Forwarded(t *testing.T) {
	target := "http://hbg.com/get"
	r := httptest.NewRequest("GET", target, nil)

	headers := make(map[string][]string)
	forwardedFor := "1.1.1.1"
	headers["X-Forwarded-For"] = []string{"1.1.1.1"}
	r.Header = headers

	hbr, err := parseRequest(r)
	if err != nil {
		t.Errorf("Failed to ParseRequest. Err: %v", err)
	}

	if hbr.Origin != forwardedFor {
		t.Errorf("got %s, want %s", hbr.Origin, forwardedFor)
	}

}

func TestParseRequest_Args(t *testing.T) {
	target := "http://hbg.com/get?test=test1,test2&Something=1"
	r := httptest.NewRequest("GET", target, nil)

	hbr, err := parseRequest(r)
	if err != nil {
		t.Errorf("Failed to ParseRequest. Err: %v", err)
	}

	if hbr.Args["test"] != "test1,test2" {
		t.Errorf("got %s, want %s", hbr.Args["test"], "test1,test2")
	}
	if hbr.Args["Something"] != "1" {
		t.Errorf("got %s, want %s", hbr.Args["Something"], "1")
	}
}

func TestToJSON(t *testing.T) {
	req := &Request{
		Args:    map[string]string{"test": "test,again"},
		Data:    "test",
		Headers: map[string]string{"Accept": "*/*", "Animal": "Dog"},
		URL:     "http://test.example.com/get?test=test,again",
		JSON:    "",
	}

	requestedKeys := req.selectKeys(requestKeys{"args", "data", "headers", "url", "json"})

	reqJSON, err := toJSON(requestedKeys)
	if err != nil {
		t.Errorf("Unable to marshal Request to JSON. Err: %v", err)
	}

	jsonParsed, err := jsonparser.ParseJSON(reqJSON)
	if err != nil {
		t.Errorf("Unable to parse returned JSON. Err: %v", err)
	}

	testCases := []struct {
		jsonPath string
		original string
	}{
		{"args.test", req.Args["test"]},
		{"data", req.Data},
		{"headers.Accept", req.Headers["Accept"]},
		{"headers.Something", req.Headers["Something"]},
		{"url", req.URL},
		{"json", req.JSON},
	}

	for _, tc := range testCases {
		if val := jsonParsed.Path(tc.jsonPath).String(); val != tc.original {
			t.Errorf("Incorrect val after unmarshal; expected: %v, got: %v", tc.original, val)
		}
	}
}

func TestRequestToJSON(t *testing.T) {
	target := "http://hbg.com/get?test=test1,test2&Something=1"
	r := httptest.NewRequest("GET", target, nil)

	requestedKeys := requestKeys{"args", "url"}

	reqJSON, err := RequestToJSON(r, requestedKeys)
	if err != nil {
		t.Errorf("Unable to marshal Request to JSON. Err: %v", err)
	}

	jsonParsed, err := jsonparser.ParseJSON(reqJSON)
	if err != nil {
		t.Errorf("Unable to parse returned JSON. Err: %v", err)
	}

	testCases := []struct {
		jsonPath string
		expected string
	}{
		{"args.test", "test1,test2"},
		{"args.Something", "1"},
		{"url", target},
	}

	for _, tc := range testCases {
		if val := jsonParsed.Path(tc.jsonPath).String(); val != tc.expected {
			t.Errorf("Incorrect val after unmarshal; expected: %v, got: %v", tc.expected, val)
		}
	}
}

func TestSelectKeys(t *testing.T) {
	target := "http://hbg.com/get?test=test1,test2&Something=1"
	r := httptest.NewRequest("GET", target, nil)

	requestedKeys := requestKeys{"args"}

	reqJSON, err := RequestToJSON(r, requestedKeys)
	if err != nil {
		t.Errorf("Unable to marshal Request to JSON. Err: %v", err)
	}

	jsonParsed, err := jsonparser.ParseJSON(reqJSON)
	if err != nil {
		t.Errorf("Unable to parse returned JSON. Err: %v", err)
	}

	testCases := []struct {
		jsonPath string
		expected string
	}{
		{"args.test", "test1,test2"},
		{"args.Something", "1"},
		{"url", ""},
		{"origin", ""},
		{"files", ""},
		{"headers", ""},
	}

	for _, tc := range testCases {
		if val := jsonParsed.Path(tc.jsonPath).String(); val != tc.expected {
			t.Errorf("Incorrect val after unmarshal; expected: %v, got: %v", tc.expected, val)
		}
	}
}
