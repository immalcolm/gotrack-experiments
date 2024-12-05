package main

/*
In this example, the program makes an HTTP GET request to the specified URL,
handles any errors that occur, and decodes the JSON response into a map.

The decoded JSON data is then printed to the console.
This approach is useful for working with APIs that return JSON data,
allowing you to easily parse and use the data within your Go program.
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	const ACCESS_KEY = "2ac7bbff61b3a079bc1a9e1675a827cc"
	url := "https://data.fixer.io/api/latest?access_key=" + ACCESS_KEY

	//using net/http package to make a GET request
	//returns 2 values: the response obj and an error
	if resp, err := http.Get(url); err != nil {
		fmt.Println(err)
	} else {
		//ensure the response body is closed after the function completes
		defer resp.Body.Close()

		//decode the response body into a map
		//map[string]interface{} is used to store the JSON data
		//type assertion is needed to extract the data
		//because we aren't sure about the type of the data
		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Println(err)
		} else {
			log.Fatal(err)
		}

		fmt.Println(result)
	}
}
