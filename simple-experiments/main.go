package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	// Declare a variable of type string
	name := "John Doe"

	// Prints "John" because it starts at index 0 and ends at index 4
	// slice of string from index n to m-1 | S[n:m]
	fmt.Println(name[0:4])
}
