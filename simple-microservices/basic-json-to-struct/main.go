package main

import (
	"encoding/json"
	"fmt"
)

// basic struct demo
type People struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Fruit struct {
	Fruit   string `json:"fruit"`
	Size    string `json:"size"`
	Color   string `json:"color"`
	special string `json:"special"` //this field will not be unmarshalled
}

func main() {
	printPeople()
	printFruit()
	unmarshalToMap()
	unmarshalToArray()
}

func printPeople() {
	var person People

	jsonString := `{"firstname": "John", "lastname": "Doe"}`

	//if there's error in the json string, it will be stored in the err variable
	err := json.Unmarshal([]byte(jsonString), &person)

	fmt.Println(person.Firstname)
	fmt.Println(person.Lastname)

	fmt.Println(err)
}

func printFruit() {
	var fruit Fruit

	jsonString := `{"fruit": "Apple", "size": "Medium", "color": "Red", "special": "Special"}`

	err := json.Unmarshal([]byte(jsonString), &fruit)

	fmt.Println(fruit.Fruit)
	fmt.Println(fruit.Size)
	fmt.Println(fruit.Color)

	fmt.Println(err)
}

func unmarshalToArray() {
	var persons []People //People struct
	jsonString :=
		`[
		{"firstname": "John", "lastname": "Doe"},
		{"firstname": "Jane", "lastname": "Doe"}
	]`

	err := json.Unmarshal([]byte(jsonString), &persons)

	fmt.Println("Unmarshalled data----------------")
	fmt.Println(persons)

	fmt.Println(err)
}

func unmarshalToMap() {
	//interface offers flexibility different types of data
	//map is a collection of key-value pairs
	var result map[string]interface{}
	jsonString := `{"fruit": "Apple", "size": "Medium", "color": "Red"}`

	//result used to store the unmarshalled data
	err := json.Unmarshal([]byte(jsonString), &result)

	fmt.Println(result)

	fmt.Println(err)
}
