package httpbin

import "net/http"

func (s *Server) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"args", "data", "files", "form", "headers", "json", "origin", "url"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"url", "args", "headers", "origin"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handlePatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"args", "data", "files", "form", "headers", "json", "origin", "url"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handlePut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"args", "data", "files", "form", "headers", "json", "origin", "url"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handlePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"args", "data", "files", "form", "headers", "json", "origin", "url"}
		returnRequestAsJSON(w, r, keys)
	}
}
