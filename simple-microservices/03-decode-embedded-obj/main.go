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

func main() {
	unmarshalEmbedded()
}

func unmarshalEmbedded() {
	var persons []People

	jsonString := `[
		{
			"firstname": "John",
			"lastname": "Doe",
			"details": {
				"age": 25,
				"occupation": "Software Engineer",
				"address": "48 Clementi Street"
			}
		},{
			"firstname": "Jane",
			"lastname": "Doe",
			"details": {
				"age": 30,
				"occupation": "Doctor",
				"address": "1234 North Bridge Road"
			}
		}
	]`

	err := json.Unmarshal([]byte(jsonString), &persons)

	if err != nil {
		fmt.Println("Error unmarshalling JSON: %v\n", err)
	}

	for _, v := range persons {
		fmt.Println(v.Firstname)
		fmt.Println(v.Lastname)
		fmt.Println(v.Details.Age)
		fmt.Println(v.Details.Occupation)
		fmt.Println(v.Details.Address)
		fmt.Printf("Person details are: %s %s %d %s %s\n--------\n",
			v.Firstname,
			v.Lastname,
			v.Details.Age,
			v.Details.Occupation,
			v.Details.Address)
	}
	fmt.Printf("=== End of Program ===\n---------------------------------")

}
