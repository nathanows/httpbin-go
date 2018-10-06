package httpbin

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var envCookies = []string{
	"_gauges_unique",
	"_gauges_unique_year",
	"_gauges_unique_month",
	"_gauges_unique_day",
	"_gauges_unique_hour",
	"__utmz",
	"__utma",
	"__utmb",
}

func (s *Server) handleCookies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var showEnv bool
		if _, ok := r.URL.Query()["show_env"]; ok {
			showEnv = true
		}
		cookies := make(map[string]string)
		for _, cookie := range r.Cookies() {
			if showEnv || !stringInSlice(cookie.Name, envCookies) {
				cookies[cookie.Name] = cookie.Value
			}
		}
		response := map[string]map[string]string{"cookies": cookies}
		json, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json = append(json, "\n"...)
		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	}
}

func (s *Server) handleCookiesDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var toDelete []string
		for k := range r.URL.Query() {
			toDelete = append(toDelete, k)
		}
		for _, cookie := range r.Cookies() {
			if stringInSlice(cookie.Name, toDelete) {
				c := http.Cookie{
					Name:    cookie.Name,
					Path:    "/",
					Value:   "asdf",
					Expires: time.Now().Add(-100 * time.Hour),
					MaxAge:  -1,
				}
				http.SetCookie(w, &c)
			}
		}
		http.Redirect(w, r, "/cookies", http.StatusFound)
	}
}

func (s *Server) handleCookiesSet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars["name"] != "" && vars["value"] != "" {
			c := http.Cookie{
				Name:    vars["name"],
				Path:    "/",
				Value:   vars["value"],
				Expires: time.Now().Add(3200 * time.Second),
				MaxAge:  3200,
			}
			http.SetCookie(w, &c)

		}

		for k, v := range r.URL.Query() {
			c := http.Cookie{
				Name:    k,
				Path:    "/",
				Value:   strings.Join(v, ","),
				Expires: time.Now().Add(3200 * time.Second),
				MaxAge:  3200,
			}
			http.SetCookie(w, &c)
		}

		http.Redirect(w, r, "/cookies", http.StatusFound)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
