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

func allcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of all courses")

	// returns the key/value pairs in the query string as a map object
	kv := r.URL.Query()

	/*
		map[string][]string{
			"country":    {"US"},
			"state": {"CA"},
		}
	*/

	for k, v := range kv {
		fmt.Println(k, v) // print out the key/value pair
	}

}

func course(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "Detail for course "+params["courseid"])
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/courses", allcourses)
	router.HandleFunc("/api/v1/courses/{courseid}", course)

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
