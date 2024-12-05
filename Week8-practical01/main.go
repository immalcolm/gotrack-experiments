package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string //first char must be uppercase | exported field
	Lastname  string //first char must be uppercase | exported field
}

func main() {
	var person People

	jsonString := `{"Firstname":"John", "Lastname":"Doe"}`

	// Unmarshal the jsonString into the person map
	// The Unmarshal function takes a byte slice of the JSON string and
	// a reference to the map where the data will be stored.
	err := json.Unmarshal([]byte(jsonString), &person) //taps on the encoding/json package

	fmt.Println(person.Firstname)
	fmt.Println(person.Lastname)
	fmt.Println(err)
}

/*

+---------------------+
|  Start of main()    |
+---------------------+
          |
          v
+---------------------+
| Declare person      |
| of type People      |
+---------------------+
          |
          v
+---------------------+
| Define jsonString   |
| with JSON data      |
+---------------------+
          |
          v
+---------------------+
| Call json.Unmarshal |
| with jsonString and |
| &person             |
+---------------------+
          |
          v
+---------------------+
| Unmarshal JSON data |
| into person struct  |
+---------------------+
          |
          v
+---------------------+
| Check for error     |
| in err variable     |
+---------------------+
          |
          v
+---------------------+
| Print person fields |
| Firstname and       |
| Lastname            |
+---------------------+
          |
          v
+---------------------+
| Print err variable  |
+---------------------+
          |
          v
+---------------------+
|  End of main()      |
+---------------------+
*/
