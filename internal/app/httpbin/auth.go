package httpbin

import (
	"crypto/subtle"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type authResponse struct {
	Authenticated bool   `json:"authenticated"`
	Token         string `json:"token,omitempty"`
	User          string `json:"user,omitempty"`
}

func (s *Server) handleBasicAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["user"]
		password := vars["password"]

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="Fake Realm"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
			return
		}

		resp := authResponse{Authenticated: true, User: username}
		jsonResp, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResp = append(jsonResp, "\n"...)

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}

func (s *Server) handleBearer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var token string

		tokens, ok := r.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
		}

		if token == "" || token == tokens[0] {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		resp := authResponse{Authenticated: true, Token: token}
		jsonResp, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResp = append(jsonResp, "\n"...)

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}
