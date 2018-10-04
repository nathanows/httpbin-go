package httpbin

import "net/http"

func (s *Server) handleAnything() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"args", "data", "files", "form", "headers", "json", "method", "origin", "url"}
		returnRequestAsJSON(w, r, keys)
	}
}
