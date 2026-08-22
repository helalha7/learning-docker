package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go Demo Server!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Handler)

	log.Println("Starting server on port :8080 ...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
