package httpbin

import (
	"net/http"
	"strconv"
)

func (s *Server) handleRedirectTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode := http.StatusFound
		if code := r.URL.Query().Get("status_code"); code != "" {
			code, err := strconv.Atoi(code)
			if err == nil && code >= 300 && code < 400 {
				statusCode = code
			}
		}
		url := r.URL.Query().Get("url")

		http.Redirect(w, r, url, statusCode)
	}
}
