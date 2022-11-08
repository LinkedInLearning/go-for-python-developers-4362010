package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("error: %s", err)
	}
}
