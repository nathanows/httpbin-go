package httpbin

import (
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) handleBase64Decode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := mux.Vars(r)["value"]
		decoded, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(decoded))
	}
}

func (s *Server) handleBytes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		length, err := parseURLFloat(mux.Vars(r)["n"], "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		length = math.Min(length, 100*1024) // set 100KB limit

		if seed := r.URL.Query().Get("seed"); seed != "" {
			i, err := strconv.Atoi(seed)
			if err == nil {
				rand.Seed(int64(i))
			}
		} else {
			rand.Seed(time.Now().UnixNano())
		}

		w.Header().Add("Content-Type", "application/octet-stream")

		randBytes := make([]byte, int(length))
		rand.Read(randBytes)
		w.Write(randBytes)
	}
}

func (s *Server) handleDelay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delay, err := parseURLFloat(mux.Vars(r)["delay"], "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := delayRequest(delay); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		keys := requestKeys{"url", "args", "form", "data", "origin", "headers", "files"}
		returnRequestAsJSON(w, r, keys)
	}
}

func (s *Server) handleDrip() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		var code string
		var err error
		var duration, numbytes, delay float64
		if duration, err = parseURLFloat(query.Get("duration"), "2"); err != nil {
			http.Error(w, "Invalid duration", http.StatusBadRequest)
			return
		}
		if numbytes, err = parseURLFloat(query.Get("numbytes"), "10"); err != nil {
			http.Error(w, "Invalid numbytes", http.StatusBadRequest)
			return
		}
		if delay, err = parseURLFloat(query.Get("delay"), "0"); err != nil {
			http.Error(w, "Invalid delay", http.StatusBadRequest)
			return
		}
		if code = query.Get("code"); code == "" {
			code = "200"
		}

		if err := delayRequest(delay); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("Content-Length", strconv.FormatFloat(numbytes, 'f', -1, 64))

		pause := duration / numbytes
		for i := 1; i <= int(numbytes); i++ {
			w.Write([]byte("*"))
			if err := delayRequest(pause); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		respCode, err := strconv.Atoi(code)
		if err != nil {
			http.Error(w, "Invalid status code provided", http.StatusBadRequest)
			return
		}
		w.WriteHeader(respCode)
	}
}

func delayRequest(delay float64) error {
	// for delays more granular than 1 sec, a delay of over 1000 will be
	// parsed as microseconds
	useMicroS := delay > 999
	if useMicroS {
		delay = math.Min(delay, 1000000) // microsecond delay must be less than 1 sec
		time.Sleep(time.Duration(delay) * time.Microsecond)
	} else {
		delay = math.Min(delay, 10) // max of 10 sec delay
		time.Sleep(time.Duration(delay) * time.Second)
	}

	return nil
}

func parseURLFloat(val, fallback string) (float64, error) {
	var parsedVal float64
	if val == "" {
		if fallback == "" {
			return 0, fmt.Errorf("no value or fallback provided")
		}
		val = fallback
	}
	parsedVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}

	if parsedVal < 0 {
		return 0, err
	}

	return parsedVal, nil
}
