package httpbin

import (
	"net/http"
	"text/template"
)

const angryASCII = `
          .-''''''-.
        .' _      _ '.
       /   O      O   \
      :                :
      |                |
      :       __       :
       \  .--"  "--.  /
        '.          .'
          '-......-'
     YOU SHOULDN'T BE HERE
`

const robotTxt = `User-agent: *
Disallow: /deny
`

func (s *Server) handleDeny() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(angryASCII))
	}
}

func (s *Server) handleEncodingUTF8() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		tmpl := template.Must(template.ParseFiles("templates/UTF-8-demo.txt"))
		tmpl.Execute(w, "")
	}
}

func (s *Server) handleHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		tmpl := template.Must(template.ParseFiles("templates/moby.html"))
		tmpl.Execute(w, "")
	}
}

func (s *Server) handleJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		jsonBlob := `{
            "title": "Sample Slide Show",
            "date": "date of publication",
            "author": "Yours Truly",
            "slides": [
                {"type": "all", "title": "Wake up to WonderWidgets!"},
                {
                    "type": "all",
                    "title": "Overview",
                    "items": [
                        "Why <em>WonderWidgets</em> are great",
                        "Who <em>buys</em> WonderWidgets"
                    ]
                }
            ]
        }`
		bytes := []byte(jsonBlob)
		w.Write(bytes)
	}
}

func (s *Server) handleRobotsTxt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(robotTxt))
	}
}

func (s *Server) handleXML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/xml")
		tmpl := template.Must(template.ParseFiles("templates/sample.xml"))
		tmpl.Execute(w, "")
	}
}
