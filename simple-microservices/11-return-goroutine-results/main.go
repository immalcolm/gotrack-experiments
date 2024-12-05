package main

/*
this Go program demonstrates how to perform concurrent HTTP requests to
different APIs, process the JSON responses, and handle the results using channels.
The use of goroutines allows the program to fetch data from multiple sources
simultaneously, making it efficient and responsive.
*/
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	apis map[int]string
	c    chan map[int]interface{} // channel to store map[int]interface{}
)

func fetchData(API int) {
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
				// store the result into the channel
				c <- re
				fmt.Println("Result for API 1 stored")

			case 2:
				if result["main"] != nil {
					re[API] = result["main"].(map[string]interface{})["temp"]
				} else {
					re[API] = result["message"]
				}
				// store the result into the channel
				c <- re
				fmt.Println("Result for API 2 stored")

			case 3:
				if result != nil {
					//get the name of the pokemon
					fmt.Println(result["name"])
					fmt.Printf("Pokemon details for %s are ,,,%s\n", result["name"], result["abilities"])
				} else {
					fmt.Println(result["error"])
				}
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func main() {
	// create a channel to store the results
	c = make(chan map[int]interface{})

	// create a map to store the APIs
	apis = make(map[int]string)

	// define the APIs
	apis[1] = "http://data.fixer.io/api/latest?access_key=<access_key>"
	apis[2] =
		"http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=<app_id>"
	apis[3] = "https://pokeapi.co/api/v2/pokemon/ditto"

	// Starting Goroutines: The go keyword is used to start new goroutines.
	go fetchData(1)
	go fetchData(2)
	go fetchData(3)

	// we expect two results in the channel
	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done!")

	fmt.Scanln()
}
