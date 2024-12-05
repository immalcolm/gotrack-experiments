package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}

func main() {
	// Create a new router
	// The router is the main router for your application.
	// It will match incoming requests against its list of routes and
	// call the handler for the route that matches the URL or other conditions.
	router := mux.NewRouter()

	//define a route and a handler function
	router.HandleFunc("/api/v1/", home)

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
