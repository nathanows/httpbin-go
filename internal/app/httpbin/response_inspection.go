package httpbin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

func (s *Server) handleCache() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("If-Modified-Since") != "" || r.Header.Get("If-None-Match") != "" {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		lastMod := time.Now().Format("Mon, 04 Oct 2018 23:08:16 GMT")
		w.Header().Set("Last-Modified", lastMod)

		uuid, _ := uuid.NewV4()
		w.Header().Set("ETag", uuid.String())

		keys := requestKeys{"args", "headers", "origin", "url"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handleCacheControl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		value := vars["value"]

		cacheControlVal := fmt.Sprintf("public, max-age=%s", value)
		w.Header().Set("Cache-Control", cacheControlVal)

		keys := requestKeys{"args", "headers", "origin", "url"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handleETag() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		etag := vars["etag"]

		ifNoneMatch := r.Header.Get("If-None-Match")
		ifNoneMatchSlice := strings.Split(ifNoneMatch, ", ")
		ifMatch := r.Header.Get("If-Match")
		ifMatchSlice := strings.Split(ifMatch, ", ")

		if ifNoneMatch != "" {
			if sliceContains(ifNoneMatchSlice, etag) || sliceContains(ifNoneMatchSlice, "*") {
				w.Header().Set("ETag", etag)
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		if ifMatch != "" {
			if sliceContains(ifMatchSlice, etag) || sliceContains(ifMatchSlice, "*") {
				w.WriteHeader(http.StatusPreconditionFailed)
				return
			}
		}

		w.Header().Set("ETag", etag)
		keys := requestKeys{"args", "headers", "origin", "url"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handleResponseHeaders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.URL.Query() {
			w.Header()[k] = v
		}

		w.Header().Set("Content-Type", "application/json")

		resp := map[string]string{}
		for k, v := range w.Header() {
			resp[k] = strings.Join(v, ",")
		}
		fmt.Printf("%+v\n", resp)

		json, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
