package main

import (
	"log"
	"net/http"
)

func main() {
	appl := newApplication()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ad/{id}", appl.getAdvertismentHandler)
	mux.HandleFunc("POST /ad/", appl.createAdvertismentHandler)

	addr := ":8080"
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
