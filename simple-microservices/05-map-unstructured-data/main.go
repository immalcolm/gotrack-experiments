package main

import (
	"encoding/json"
	"fmt"
)

type Rates struct {
	Base   string `json:"base currency"` //allow more control, the json will be "base currency"
	Symbol string `json:"destination currency"`
}

func main() {

}

func mapUnstructuredData() {
	jsonString := `{
		"success": true,
		"timestamp": 1626591606,
		"base": "USD",
		"date": "2024-12-05",
		"rates": {	
			"USD": 1,
			"SGD": 1.36,
			"EUR": 0.82,
			"JPY": 110.2
		}
	}`

	//because the json structure is not fixed,
	//we can use map[string]interface{}
	//type assertion is needed to extract the data
	//which is a map with string keys and values of any type.
	var result map[string]interface{}

	json.Unmarshal([]byte(jsonString), &result)

	//we can access the data using the key
	fmt.Println(result["success"])

	fmt.Println(result["rates"])

	//type assertion and map[string]interface{}
	//extract a specific value from the nested map structure
	SGD := result["rates"].(map[string]interface{})["SGD"]

	fmt.Println("SGD to USD rate is", SGD)

}
