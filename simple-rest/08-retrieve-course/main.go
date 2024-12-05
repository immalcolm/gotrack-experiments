package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type courseInfo struct {
	Title string `json:"Title"`
}

// used for storing courses on the REST API
var courses map[string]courseInfo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}

func allcourses(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "List of all courses")

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

	//returns all the courses in JSON format
	json.NewEncoder(w).Encode(courses)
}

// function to handle course requests
// GET, PUT, POST, DELETE
// only POST is implemented
// POST is used to create a new course
// Sample query: http://localhost:5001/api/v1/courses/T01
// with JSON body {"Title":"Math 101"}
func course(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	/*
		fmt.Fprintf(w, "Detail for course "+params["courseid"])

		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, r.Method)
	*/

	// [MY] GET is for retrieving course information
	if r.Method == "GET" {
		if _, ok := courses[params["courseid"]]; ok {
			json.NewEncoder(w).Encode(
				courses[params["courseid"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new course
		if r.Method == "POST" {

			// read the string sent to the service
			var newCourse courseInfo
			reqBody, err := io.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newCourse)

				if newCourse.Title == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply course " +
							"information " + "in JSON format"))
					return
				}

				// check if course exists; add only if
				// course does not exist
				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " +
						params["courseid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate course ID"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information " +
					"in JSON format"))
			}
		}

		//---PUT is for creating or updating
		// existing course---
		if r.Method == "PUT" {
			var newCourse courseInfo
			reqBody, err := io.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newCourse)

				if newCourse.Title == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply course " +
							" information " +
							"in JSON format"))
					return
				}

				// check if course exists; add only if
				// course does not exist
				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] =
						newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " +
						params["courseid"]))
				} else {
					// update course
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Course updated: " +
						params["courseid"]))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " +
					"course information " +
					"in JSON format"))
			}
		}

	}

}

func main() {

	// initialize the courses map
	courses = make(map[string]courseInfo)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/courses", allcourses)

	//You can register each HTTP method with individual functions,
	//or use a single function
	//if method is not listed, it will return 405 Method Not Allowed
	router.HandleFunc("/api/v1/courses/{courseid}", course).Methods(
		"GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
