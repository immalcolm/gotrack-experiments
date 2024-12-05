package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	apis map[int]string
	c    chan map[int]interface{} // channel to store map[int]interface{}
	wg   sync.WaitGroup           // WaitGroup to wait for all goroutines to finish
)

func fetchData(API int) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {

			var result map[string]interface{}
			json.Unmarshal(body, &result)

			var re = make(map[int]interface{})

			switch API {
			case 1:
				if result["success"] == true {
					re[API] = result["rates"].(map[string]interface{})["USD"]
				} else {
					re[API] = result["error"].(map[string]interface{})["info"]
				}
			case 2:
				if result["cod"] == 200.0 {
					re[API] = result["main"].(map[string]interface{})["temp"]
				} else {
					re[API] = result["message"]
				}
			case 3:
				if result != nil {
					re[API] = result["name"]
				} else {
					re[API] = result["error"]
				}
			}

			// Store the result into the channel
			c <- re
		}
	}
}

func main() {
	apis = make(map[int]string)
	apis[1] = "http://data.fixer.io/api/latest?access_key=2ac7bbff61b3a079bc1a9e1675a827cc"
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?q=Singapore&appid=8ea49a70241b4c7ce4839766bcd9f267"
	apis[3] = "https://pokeapi.co/api/v2/pokemon/ditto"

	c = make(chan map[int]interface{}, 3) // Buffered channel with capacity 3

	// Add a counter to the WaitGroup for each goroutine
	wg.Add(3)

	go fetchData(1)
	go fetchData(2)
	go fetchData(3)

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channel after all goroutines have finished
	close(c)

	// Read and print results from the channel
	for result := range c {
		fmt.Println(result)
	}

	fmt.Println("All tasks completed.")
}
