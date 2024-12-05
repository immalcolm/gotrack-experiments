package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	unmarshalArray()
}

func unmarshalArray() {
	var persons []People
	jsonString := `[
		{"firstname": "John", "lastname": "Doe"},
		{"firstname": "Jane", "lastname": "Doe"},
		{"firstname": "Jamie", "lastname": "Oliver"}
	]`

	err := json.Unmarshal([]byte(jsonString), &persons)

	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
	}

	for _, v := range persons {
		fmt.Println(v.Firstname)
		fmt.Println(v.Lastname)
		fmt.Printf("Person details are: %s %s\n--------\n", v.Firstname, v.Lastname)
	}
}
