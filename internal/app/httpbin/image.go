package httpbin

import (
	"net/http"
	"regexp"
)

var images = map[string]string{
	"image/jpeg":    "images/jackal.jpg",
	"image/png":     "images/pig_icon.png",
	"image/webp":    "images/wolf_1.webp",
	"image/svg+xml": "images/svg_logo.svg",
}

func (s *Server) handleImage(imageType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imgType := imageType
		if imgType == "" {
			re := regexp.MustCompile("image/[a-z\\+]+")
			imgType = re.FindString(r.Header.Get("accept"))
			if imgType == "" {
				imgType = "image/png"
			}
		}
		imgPath, ok := images[imgType]
		if !ok {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		w.Header().Set("Content-Type", imgType)
		http.ServeFile(w, r, imgPath)
	}
}
