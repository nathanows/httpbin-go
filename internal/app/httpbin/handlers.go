package httpbin

import (
	"net/http"
)

func returnRequestAsJSON(w http.ResponseWriter, r *http.Request, keys requestKeys) {
	json, err := RequestToJSON(r, keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
