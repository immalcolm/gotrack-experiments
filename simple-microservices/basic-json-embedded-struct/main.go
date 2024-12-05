package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Details   struct {
		Age        int    `json:"age"`
		Occupation string `json:"occupation"`
	}
}

func main() {
	unmarshalEmbedded()
}

func unmarshalEmbedded() {

	var person []People

	jsonString := `[
		{
			"firstname": "John",
			"lastname": "Doe",
			"details": {
				"age":25,
				"occupation":"Software Engineer"
			}
		},
		{
			"firstname": "Jane",
			"lastname": "Doe",
			"details": {
				"age": 30,
				"occupation": "Doctor"
			}
		}
	]`

	err := json.Unmarshal([]byte(jsonString), &person)

	fmt.Println(person)
	fmt.Println(err)

	for _, v := range person {
		fmt.Println(v.Firstname)
		fmt.Println(v.Lastname)
		fmt.Println(v.Details.Age)
		fmt.Println(v.Details.Occupation)
		fmt.Printf("Person details are: %s %s %d %s\n--------\n",
			v.Firstname,
			v.Lastname,
			v.Details.Age,
			v.Details.Occupation)
	}

}
