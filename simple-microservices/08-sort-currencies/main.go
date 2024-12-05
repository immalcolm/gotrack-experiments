package main

/*
In this example, the program makes an HTTP GET request to the
specified URL, handles the response, and processes the JSON data.

If the request is successful and the JSON data indicates success,
it prints the exchange rates in alphabetical order.
If there is an error, it prints the error information.

*/
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

type Result struct {
	Success   bool               `json:"success"`
	Timestamp int                `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

type Error struct {
	Success bool `json:"success"`
	Error   struct {
		Code int    `json:"code"`
		Type string `json:"type"`
		Info string `json:"info"`
	}
}

func main() {
	const APIKEY = "2ac7bbff61b3a079bc1a9e1675a827cc"
	url := "http://data.fixer.io/api/latest?access_key=" + APIKEY

	if resp, err := http.Get(url); err == nil {
		//ensure the response body is closed after the function completes
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			//dump to result
			var result Result

			//unmarshall the JSON data into the result struct
			json.Unmarshal(body, &result)

			if result.Success {
				//do something with the data
				//create a slice to store the keys
				keys := make([]string, 0, len(result.Rates))

				//get all the keys
				for k := range result.Rates {
					keys = append(keys, k)
				}

				sort.Strings(keys)

				//print in alphabetical order
				for _, k := range keys {
					fmt.Println(k, result.Rates[k])
				}

				/*
					for i, v := range result.Rates {
						fmt.Println(i, v)
					}
				*/
			} else {
				//unmarshall the content to error
				var err Error
				json.Unmarshal(body, &err)

				fmt.Println(err.Error.Info)
			}
		} else {
			fmt.Println(err)
			log.Fatal(err)
		}
	} else {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("done----------------")
}
