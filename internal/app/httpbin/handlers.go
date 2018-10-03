package httpbin

import (
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

func returnRequestAsJSON(w http.ResponseWriter, r *http.Request, keys requestKeys) {
	json, err := RequestToJSON(r, keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
