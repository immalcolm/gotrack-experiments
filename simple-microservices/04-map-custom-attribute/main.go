package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Details   struct {
		Age        int    `json: "age"`
		Occupation string `json: "occupation"`
		Address    string `json: "address"`
	}
}

type Rates struct {
	Base   string `json:"base currency"` //allow more control, the json will be "base currency"
	Symbol string `json:"destination currency"`
}

func main() {
	mapCustomAttribute()
}

func mapCustomAttribute() {
	var rates Rates

	//this json structure is not so good due to whitespaces
	jsonString := `{"base currency": "USD", "destination currency": "SGD"}`
	//improved json structure
	//jsonString := `{"base_currency": "USD", "destination_currency": "SGD"}`

	err := json.Unmarshal([]byte(jsonString), &rates)

	fmt.Println(rates.Base)
	fmt.Println(rates.Symbol)

	fmt.Println(err)
}
