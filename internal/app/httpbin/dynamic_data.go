package httpbin

import (
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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

		respCode, err := strconv.Atoi(code)
		if err != nil {
			http.Error(w, "Invalid status code provided", http.StatusBadRequest)
			return
		}
		w.WriteHeader(respCode)

		pause := duration / numbytes
		for i := 1; i <= int(numbytes); i++ {
			w.Write([]byte("*"))
			if err := delayRequest(pause); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func (s *Server) handleLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var err error
		var n, offset float64
		if n, err = parseURLFloat(vars["n"], ""); err != nil {
			http.Error(w, "Invalid number of links", http.StatusBadRequest)
			return
		}
		n = math.Min(n, 200) // max of 200 links
		if offset, err = parseURLFloat(vars["offset"], ""); err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}

		html := []string{"<html><head><title>Links</title></head><body>"}

		for i := 0; i < int(n); i++ {
			if i == int(offset) {
				html = append(html, fmt.Sprintf("%d ", i))
			} else {
				html = append(html, fmt.Sprintf("<a href='/links/%d/%d'>%d</a> ", int(n), i, i))
			}
		}

		html = append(html, "</body></html>")
		w.Write([]byte(strings.Join(html, "")))
	}
}

func (s *Server) handleRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		numbytes, err := parseURLFloat(mux.Vars(r)["numbytes"], "")
		if err != nil {
			http.Error(w, "Invalid numbytes", http.StatusBadRequest)
			return
		}

		if numbytes > (100 * 1024) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("number of bytes must be in the range (0, 102400)"))
			return
		}

		var chunkSize float64
		if chunkSize, err = parseURLFloat(r.URL.Query().Get("chunk_size"), "10240"); err != nil {
			http.Error(w, "Invalid chunk_size", http.StatusBadRequest)
			return
		}

		var duration float64
		if duration, err = parseURLFloat(r.URL.Query().Get("duration"), "0"); err != nil {
			http.Error(w, "Invalid duration", http.StatusBadRequest)
			return
		}

		pausePerByte := duration / numbytes
		firstBytePos, lastBytePos := getRequestRange(r.Header, int(numbytes))
		rangeLength := (lastBytePos + 1) - firstBytePos

		if firstBytePos > lastBytePos || !inRange(firstBytePos, 0, int(numbytes)) || !inRange(lastBytePos, 0, int(numbytes)) {
			w.Header().Set("ETag", fmt.Sprintf("range%d", int(numbytes)))
			w.Header().Set("Accept-Ranges", "bytes")
			w.Header().Set("Content-Length", "0")
			w.Header().Set("Content-Range", fmt.Sprintf("bytes */%d", int(numbytes)))
			w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
			return
		}

		contentRange := fmt.Sprintf("bytes %d-%d/%g", firstBytePos, lastBytePos, numbytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("ETag", fmt.Sprintf("range%d", int(numbytes)))
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Content-Range", contentRange)
		w.Header().Set("Content-Length", strconv.Itoa(rangeLength))

		if firstBytePos == 0 && int64(lastBytePos) == int64(numbytes-1) {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusPartialContent)
		}

		fw := flushWriter{w: w}
		if f, ok := w.(http.Flusher); ok {
			fw.f = f
		}

		var chunks []rune

		for i := firstBytePos; i <= lastBytePos; i++ {
			chunk := int('a') + (i % 26)
			chunks = append(chunks, rune(chunk))
			if len(chunks) == int(chunkSize) {
				fw.Write([]byte(string(chunks)))
				if err := delayRequest(pausePerByte * chunkSize); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				chunks = []rune{}
			}
		}

		if len(chunks) > 0 {
			lastDelay, err := strconv.ParseFloat(strconv.Itoa(len(chunks)), 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}
			lastDelay = pausePerByte * lastDelay
			if err := delayRequest(lastDelay); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fw.Write([]byte(string(chunks)))
		}
	}
}

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

func getRequestRange(reqHeaders http.Header, upperBound int) (firstBytePos, lastBytePos int) {
	firstBytePos, lastBytePos = parseRequestRange(reqHeaders.Get("range"))

	if firstBytePos == -1 && lastBytePos == -1 {
		firstBytePos = 0
		lastBytePos = upperBound - 1
	} else if firstBytePos == -1 {
		if 0 > (upperBound - lastBytePos) {
			firstBytePos = 0
		}
		firstBytePos = upperBound - lastBytePos
		lastBytePos = upperBound - 1
	} else if lastBytePos == -1 {
		lastBytePos = upperBound - 1
	}

	return firstBytePos, lastBytePos
}

// Return a tuple describing the byte range requested in a GET request
// If the range is open ended on the left or right side, then a value of None
// will be set.
// RFC7233: http://svn.tools.ietf.org/svn/wg/httpbis/specs/rfc7233.html#header.range
// Examples:
//   Range : bytes=1024-
//   Range : bytes=10-20
//   Range : bytes=-999
func parseRequestRange(rangeHeaderText string) (left, right int) {
	var err error
	rawRangeHeader := strings.TrimSpace(rangeHeaderText)
	if !strings.HasPrefix(rawRangeHeader, "bytes") {
		return -1, -1
	}

	components := strings.Split(rawRangeHeader, "=")
	if len(components) != 2 {
		return -1, -1
	}

	components = strings.Split(components[1], "-")
	left, err = strconv.Atoi(components[0])
	if err != nil {
		return -1, -1
	}
	right, err = strconv.Atoi(components[1])
	if err != nil {
		return -1, -1
	}

	return left, right
}

func inRange(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	}
	return false
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
