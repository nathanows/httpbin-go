package httpbin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ndwhtlssthr/httpbin-go/pkg/jsonparser"
)

var emptyServer = &Server{}

func TestHandleDelete(t *testing.T) {
	target := "http://test.com/delete?something=good"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	jsonParsed, err := makeRequest(emptyServer.handleDelete(), target, "DELETE", headers)
	if err != nil {
		t.Errorf("Failed to make and parse request. Err: %v", err)
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

	validateCorrectFields(t, []string{"args", "data", "files", "form", "headers", "json", "origin", "url"}, jsonParsed)
}

func TestHandleGet(t *testing.T) {
	target := "http://test.com/get?something=else"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	jsonParsed, err := makeRequest(emptyServer.handleGet(), target, "GET", headers)
	if err != nil {
		t.Errorf("Failed to make and parse request. Err: %v", err)
	}

	fmt.Println("here in GET")
	fmt.Printf("%#v\n", jsonParsed)

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

	validateCorrectFields(t, []string{"args", "headers", "origin", "url"}, jsonParsed)
}

func TestHandlePatch(t *testing.T) {
	target := "http://test.com/patch?something=patched"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	jsonParsed, err := makeRequest(emptyServer.handlePatch(), target, "PATCH", headers)
	if err != nil {
		t.Errorf("Failed to make and parse request. Err: %v", err)
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

	validateCorrectFields(t, []string{"args", "data", "files", "form", "headers", "json", "origin", "url"}, jsonParsed)
}

func TestHandlePut(t *testing.T) {
	target := "http://test.com/put?something=put"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	jsonParsed, err := makeRequest(emptyServer.handlePut(), target, "PUT", headers)
	if err != nil {
		t.Errorf("Failed to make and parse request. Err: %v", err)
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

	validateCorrectFields(t, []string{"args", "data", "files", "form", "headers", "json", "origin", "url"}, jsonParsed)
}

func TestHandlePost(t *testing.T) {
	target := "http://test.com/post?something=post"
	headers := map[string][]string{"Accept": []string{"*/*"}}
	jsonParsed, err := makeRequest(emptyServer.handlePost(), target, "POST", headers)
	if err != nil {
		t.Errorf("Failed to make and parse request. Err: %v", err)
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

	validateCorrectFields(t, []string{"args", "data", "files", "form", "headers", "json", "origin", "url"}, jsonParsed)
}

func makeRequest(handlerFunc http.HandlerFunc, target, method string, headers map[string][]string) (*jsonparser.Container, error) {
	req, err := http.NewRequest(method, target, nil)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	rr := httptest.NewRecorder()

	handlerFunc.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		return nil, fmt.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	jsonParsed, err := jsonparser.ParseJSON(rr.Body.Bytes())
	if err != nil {
		return nil, fmt.Errorf("Unable to parse returned JSON. Err: %v", err)
	}

	return jsonParsed, nil
}

func validateCorrectFields(t *testing.T, expected []string, json *jsonparser.Container) {
	for _, field := range expected {
		if val := json.Path(field).Data(); val == nil {
			t.Errorf("Expected field %s to be included in response", field)
		}
	}

	allFields := []string{"args", "data", "files", "form", "headers", "json", "origin", "url"}
	expectedNotIncluced := sliceDiff(allFields, expected)

	for _, field := range expectedNotIncluced {
		if val := json.Path(field).Data(); val != nil {
			fmt.Printf("%#v\n", val)
			t.Errorf("%s should not be included in response, got: %s", field, json.Path(field).String())
		}
	}

}

func sliceDiff(slice1 []string, slice2 []string) []string {
	var diff []string

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
