package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", Handler)

	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
