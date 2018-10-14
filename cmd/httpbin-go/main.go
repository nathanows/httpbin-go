package main

import (
	"log"

	"github.com/nathanows/httpbin-go/internal/app/httpbin"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	server, err := httpbin.NewServer(router)
	if err != nil {
		log.Fatalf("Unable to setup server. Err: %+v", err)
	}

	server.ListenAndServe()
}
