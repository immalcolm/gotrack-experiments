package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			//dump to result
			var result Result
			json.Unmarshal(body, &result)

			if result.Success {
				//do something with the data
				for i, v := range result.Rates {
					fmt.Println(i, v)
				}
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
