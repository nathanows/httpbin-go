package httpbin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Request represents http request metadata
type Request struct {
	Args    map[string]string `json:"args"`
	Data    string            `json:"data"`
	Files   map[string]string `json:"files"`
	Form    map[string]string `json:"form"`
	Headers map[string]string `json:"headers"`
	JSON    string            `json:"json"`
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
}

// ParseRequestToJSON parses an incoming http request and returns a bytes.Buffer
// containing a properly indented, JSON formatted httpbin.Request
func ParseRequestToJSON(r *http.Request) (*bytes.Buffer, error) {
	req, err := parseRequest(r)
	if err != nil {
		return nil, err
	}

	buf, err := req.toJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to json: %v", err)
	}

	return buf, nil
}

func parseRequest(r *http.Request) (*Request, error) {
	return &Request{
		Args:    getArgs(r),
		Data:    "",
		Files:   make(map[string]string),
		Form:    make(map[string]string),
		Headers: getHeaders(r),
		JSON:    "",
		Origin:  getOrigin(r),
		URL:     getURL(r),
	}, nil
}

func getURL(r *http.Request) string {
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

func (req *Request) toJSON() (*bytes.Buffer, error) {
	response, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = json.Indent(buf, response, "", "  ")
	buf.Write([]byte("\n"))
	return buf, nil
}
