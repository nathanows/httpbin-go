package httpbin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ndwhtlssthr/httpbin-go/pkg/jsonparser"
)

func TestHandleDelete(t *testing.T) {
	target := "http://test.com/delete?something=good"
	req, err := http.NewRequest("DELETE", target, nil)
	if err != nil {
		t.Errorf("Unable to build request. Err: %v", err)
	}
	req.Header = map[string][]string{"Accept": []string{"*/*"}}

	rr := httptest.NewRecorder()

	server := &Server{}
	http.HandlerFunc(server.handleDelete()).
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

	includedFieldTestCases := []struct {
		field    string
		excluded bool
	}{
		{"args", true},
		{"data", true},
		{"files", true},
		{"form", true},
		{"headers", true},
		{"json", true},
		{"origin", true},
		{"url", true},
	}

	for _, tc := range includedFieldTestCases {
		if val := jsonParsed.Path(tc.field).Data(); (val == nil) == tc.excluded {
			t.Errorf("Expected field %s inclusion: %v, was %v", tc.field, tc.excluded, !tc.excluded)
		}
	}
}

func TestHandleGet(t *testing.T) {
	target := "http://test.com/get?something=else"
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		t.Errorf("Unable to build request. Err: %v", err)
	}
	req.Header = map[string][]string{"Accept": []string{"*/*"}}

	rr := httptest.NewRecorder()

	server := &Server{}
	http.HandlerFunc(server.handleGet()).
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
		{"args.something", "else"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	for _, tc := range testCases {
		if val := jsonParsed.Path(tc.jsonPath).String(); val != tc.expected {
			t.Errorf("Incorrect val after unmarshal; expected: %v, got: %v", tc.expected, val)
		}
	}

	includedFieldTestCases := []struct {
		field    string
		excluded bool
	}{
		{"args", true},
		{"data", false},
		{"files", false},
		{"form", false},
		{"headers", true},
		{"json", false},
		{"origin", true},
		{"url", true},
	}

	for _, tc := range includedFieldTestCases {
		if val := jsonParsed.Path(tc.field).Data(); (val == nil) == tc.excluded {
			t.Errorf("Expected field %s inclusion: %v, was %v", tc.field, tc.excluded, !tc.excluded)
		}
	}
}

func TestHandlePatch(t *testing.T) {
	target := "http://test.com/patch?something=patched"
	req, err := http.NewRequest("PATCH", target, nil)
	if err != nil {
		t.Errorf("Unable to build request. Err: %v", err)
	}
	req.Header = map[string][]string{"Accept": []string{"*/*"}}

	rr := httptest.NewRecorder()

	server := &Server{}
	http.HandlerFunc(server.handlePatch()).
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
		{"args.something", "patched"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	for _, tc := range testCases {
		if val := jsonParsed.Path(tc.jsonPath).String(); val != tc.expected {
			t.Errorf("Incorrect val after unmarshal; expected: %v, got: %v", tc.expected, val)
		}
	}

	includedFieldTestCases := []struct {
		field    string
		excluded bool
	}{
		{"args", true},
		{"data", true},
		{"files", true},
		{"form", true},
		{"headers", true},
		{"json", true},
		{"origin", true},
		{"url", true},
	}

	for _, tc := range includedFieldTestCases {
		if val := jsonParsed.Path(tc.field).Data(); (val == nil) == tc.excluded {
			t.Errorf("Expected field %s inclusion: %v, was %v", tc.field, tc.excluded, !tc.excluded)
		}
	}
}

func TestHandlePut(t *testing.T) {
	target := "http://test.com/put?something=put"
	req, err := http.NewRequest("PUT", target, nil)
	if err != nil {
		t.Errorf("Unable to build request. Err: %v", err)
	}
	req.Header = map[string][]string{"Accept": []string{"*/*"}}

	rr := httptest.NewRecorder()

	server := &Server{}
	http.HandlerFunc(server.handlePut()).
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
		{"args.something", "put"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	for _, tc := range testCases {
		if val := jsonParsed.Path(tc.jsonPath).String(); val != tc.expected {
			t.Errorf("Incorrect val after unmarshal; expected: %v, got: %v", tc.expected, val)
		}
	}

	includedFieldTestCases := []struct {
		field    string
		excluded bool
	}{
		{"args", true},
		{"data", true},
		{"files", true},
		{"form", true},
		{"headers", true},
		{"json", true},
		{"origin", true},
		{"url", true},
	}

	for _, tc := range includedFieldTestCases {
		if val := jsonParsed.Path(tc.field).Data(); (val == nil) == tc.excluded {
			t.Errorf("Expected field %s inclusion: %v, was %v", tc.field, tc.excluded, !tc.excluded)
		}
	}
}

func TestHandlePost(t *testing.T) {
	target := "http://test.com/post?something=post"
	req, err := http.NewRequest("POST", target, nil)
	if err != nil {
		t.Errorf("Unable to build request. Err: %v", err)
	}
	req.Header = map[string][]string{"Accept": []string{"*/*"}}

	rr := httptest.NewRecorder()

	server := &Server{}
	http.HandlerFunc(server.handlePost()).
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
		{"args.something", "post"},
		{"headers.Accept", "*/*"},
		{"url", target},
	}

	for _, tc := range testCases {
		if val := jsonParsed.Path(tc.jsonPath).String(); val != tc.expected {
			t.Errorf("Incorrect val after unmarshal; expected: %v, got: %v", tc.expected, val)
		}
	}

	includedFieldTestCases := []struct {
		field    string
		excluded bool
	}{
		{"args", true},
		{"data", true},
		{"files", true},
		{"form", true},
		{"headers", true},
		{"json", true},
		{"origin", true},
		{"url", true},
	}

	for _, tc := range includedFieldTestCases {
		if val := jsonParsed.Path(tc.field).Data(); (val == nil) == tc.excluded {
			t.Errorf("Expected field %s inclusion: %v, was %v", tc.field, tc.excluded, !tc.excluded)
		}
	}
}
