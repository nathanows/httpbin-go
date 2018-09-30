package httpbin

import (
	"fmt"
	"net/http"
)

func (s *Server) handleSomething() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s\n", "World")
	}
}
