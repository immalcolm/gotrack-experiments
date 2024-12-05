package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Result struct {
	Success bool
	Timestamp int
}

func main() {
	url := "https://data.fixer.io/api/latest?access_key=2ac7bbff61b3a079bc1a9e1675a827cc"

	//http.Get returns two values: the response and an error
	if response, err := http.Get(url); err == nil {
		//defer is used to close the response body after the response sucessful
		defer response.Body.Close()
		//io.ReadAll reads from the response body into a byte slice
		if body, err := io.ReadAll(response.Body); err == nil {
			fmt.Println(string(body))
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	fmt.Println("Done!")
}
