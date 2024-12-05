package main

import (
	"fmt"
	"net/http"
)

func main() {
	initServer()
}

// w - response writer
// r - request
func initServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requestedss: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":80", nil)
}
