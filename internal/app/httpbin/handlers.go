package httpbin

import (
	"crypto/subtle"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// HTTP Handlers

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

// Anything Handlers

func (s *Server) handleAnything() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := requestKeys{"args", "data", "files", "form", "headers", "json", "method", "origin", "url"}
		returnRequestAsJSON(w, r, keys)
	}
}

// Status Codes Handlers

func (s *Server) handleStatusCodes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strCodes := strings.Split(vars["codes"], ",")
		var codes []int
		for _, code := range strCodes {
			i, err := strconv.Atoi(code)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			codes = append(codes, i)
		}
		rand.Seed(time.Now().Unix())
		w.WriteHeader(codes[rand.Intn(len(codes))])
	}
}

// Request Inspection Handlers

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

// Auth Handlers

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

		type authResponse struct {
			Authenticated bool   `json:"authenticated"`
			User          string `json:"user"`
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

		type bearerResponse struct {
			Authenticated bool   `json:"authenticated"`
			Token         string `json:"token"`
		}

		resp := bearerResponse{Authenticated: true, Token: token}
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

func returnRequestAsJSON(w http.ResponseWriter, r *http.Request, keys requestKeys) {

	json, err := RequestToJSON(r, keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
