package httpbin

import (
	"fmt"
	"net/http"
)

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

func returnRequestAsJSON(w http.ResponseWriter, r *http.Request, keys requestKeys) {
	json, err := RequestToJSON(r, keys)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
