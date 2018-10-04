package httpbin

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

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
