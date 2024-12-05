package main

import (
	"encoding/json"
	"fmt"
)

/*
	Simple JSON to Struct Example
*/

// let's define a simple struct
// this struct will be used to unmarshal json data
// the struct will have two fields: Firstname and Lastname
// fields must have initial capital letter to be exported (public)
// fields that are in lowercase are private and not exported
type People struct {
	Firstname string `json:"firstname"` //first char to be capitalized
	Lastname  string `json:"lastname"`  //first char to be capitalized
}

func main() {
	printPeople()
}

func printPeople() {
	var person People
	//create our sample json
	jsonString := `{"firstname": "John", "lastname": "Doe"}` //whatever that is in the json string will be unmarshalled to the struct

	//unmarshal the json to the struct
	//if there's error in the json string, it will be stored in the err variable
	err := json.Unmarshal([]byte(jsonString), &person) //unmarshal the json string to the struct

	if err != nil {
		fmt.Println(err)
	}

	//print the values
	fmt.Println(person)

	//print nicely
	fmt.Printf("Firstname: %s, Lastname: %s\n", person.Firstname, person.Lastname)
}
