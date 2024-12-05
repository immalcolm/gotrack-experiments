package main

/*
refactor the code for json decoding
so that we can consume different webservices and decode the response

*/
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var apis map[int]string

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

	apis = make(map[int]string)
	apis[1] = "http://data.fixer.io/api/latest?access_key=2ac7bbff61b3a079bc1a9e1675a827cc"
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?q=Singapore&appid=8ea49a70241b4c7ce4839766bcd9f267"
	apis[3] = "https://pokeapi.co/api/v2/pokemon/ditto"

	fetchData(1)
	fetchData(2)
	fetchData(3)
}

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()

		if body, err := io.ReadAll(resp.Body); err == nil {

			var result map[string]interface{}

			json.Unmarshal(body, &result)

			switch API {
			case 1:
				if result["success"] == true {
					fmt.Println(result["rates"].(map[string]interface{})["USD"])
				} else {
					//fmt.Println(result["error"])
					fmt.Println(result["error"].(map[string]interface{})["info"])
				}
			case 2:
				if result["success"] == true {
					fmt.Print(result["main"].(map[string]interface{})["temp"])
				} else {
					fmt.Println(result["message"])
				}

			case 3:
				if result != nil {
					//get the name of the pokemon
					fmt.Println(result["name"])
					fmt.Printf("Pokemon details for %s are ,,,%s\n", result["name"], result["abilities"])
				} else {
					fmt.Println(result["error"])
				}
			} //end switch
		} //end of read
	} //end of get
} //end of fetchData
