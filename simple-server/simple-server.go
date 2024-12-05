package main

import (
	"fmt"
	"net/http"
)

func main() {
	initializeServer()
	// You can call other initialization functions here if needed
}

// setup simple HTTP server
func initializeServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
