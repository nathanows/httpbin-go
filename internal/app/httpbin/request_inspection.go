package httpbin

import "net/http"

func (s *Server) handleHeaders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"headers"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handleIP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"origin"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handleUserAgent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"user-agent"}
		returnRequestAsJSON(w, r, keys)
	}
}
