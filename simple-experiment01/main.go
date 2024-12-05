// 1. create a new module first
package main

//2. import packages
import (
	"fmt"
)

// 3. create a function
func main() {
	initLearn()
}

// declaring variables
func initLearn() {
	fmt.Println("Hello, World!")
	//explicit typing
	var aString string = "this is a string"
	var anInteger int = 42
	var aFloat float64 = 3.14
	//var impliedString := "this is an implied string"
	//fmt.Println(aString)

	fmt.Printf("variables %T %T %T\n", aString, anInteger, aFloat)

	var anotherString = "this is another string"
	fmt.Println(anotherString)

	//implicit typing
	//only allowed inside functions
	myString := "this is also a string"
	fmt.Println(myString)
}
