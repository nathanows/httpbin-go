package httpbin

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// Request represents http request metadata
type Request struct {
	Args      map[string]string `json:"args"`
	Data      string            `json:"data"`
	Files     map[string]string `json:"files"`
	Form      map[string]string `json:"form"`
	Headers   map[string]string `json:"headers"`
	JSON      string            `json:"json"`
	Origin    string            `json:"origin"`
	URL       string            `json:"url"`
	Method    string            `json:"method"`
	UserAgent string            `json:"user-agent"`
	Gzipped   bool              `json:"gzipped"`
}

type requestKeys []string

func returnRequestAsJSON(w http.ResponseWriter, r *http.Request, keys requestKeys) {
	json, err := RequestToJSON(r, keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func returnRequestGzipped(w http.ResponseWriter, r *http.Request, keys requestKeys) {
	resp, err := RequestToJSON(r, keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")

	gz := gzip.NewWriter(w)
	json.NewEncoder(gz).Encode(resp)
	gz.Close()
}

// RequestToJSON parses an incoming http request and returns a bytes.Buffer
// containing a properly indented, JSON formatted httpbin.Request
func RequestToJSON(r *http.Request, keys requestKeys) ([]byte, error) {
	req, err := parseRequest(r)
	if err != nil {
		return nil, err
	}

	requestedKeys := req.selectKeys(keys)

	json, err := toJSON(requestedKeys)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to json: %v", err)
	}

	return json, nil
}

func parseRequest(r *http.Request) (*Request, error) {
	return &Request{
		Args:      getArgs(r),
		Data:      "",
		Files:     make(map[string]string),
		Form:      make(map[string]string),
		Headers:   getHeaders(r),
		JSON:      "",
		Method:    getMethod(r),
		Origin:    getOrigin(r),
		URL:       getURL(r),
		UserAgent: getUserAgent(r),
	}, nil
}

func toJSON(in map[string]interface{}) ([]byte, error) {
	response, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return nil, err
	}
	response = append(response, "\n"...)

	return response, nil
}

func getURL(r *http.Request) string {
	if r.URL.IsAbs() {
		return r.URL.String()
	}

	scheme := r.URL.Scheme
	if scheme == "" {
		if r.TLS == nil {
			scheme = "http"
		} else {
			scheme = "https"
		}
	}

	url := fmt.Sprintf("%s://%s%s", scheme, r.Host, r.URL)
	return url
}

func getHeaders(r *http.Request) map[string]string {
	var headers = make(map[string]string)
	for key, vals := range r.Header {
		headers[key] = strings.Join(vals, ",")
	}
	return headers
}

func getOrigin(r *http.Request) string {
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		return forwardedFor
	}
	return r.RemoteAddr
}

func getArgs(r *http.Request) map[string]string {
	var args = make(map[string]string)
	for key, vals := range r.URL.Query() {
		args[key] = strings.Join(vals, ",")
	}
	return args
}

func getMethod(r *http.Request) string {
	return r.Method
}

func getUserAgent(r *http.Request) string {
	return r.Header.Get("User-Agent")
}

func keySet(keys requestKeys) map[string]bool {
	set := make(map[string]bool, len(keys))
	for _, s := range keys {
		set[s] = true
	}
	return set
}

func (req *Request) selectKeys(keys requestKeys) map[string]interface{} {
	ks := keySet(keys)
	rt, rv := reflect.TypeOf(*req), reflect.ValueOf(*req)
	out := make(map[string]interface{}, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		jsonKey := field.Tag.Get("json")
		if ks[jsonKey] {
			out[jsonKey] = rv.Field(i).Interface()
		}
	}
	return out
}
