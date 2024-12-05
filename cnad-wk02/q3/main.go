package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	//get user input
	var days int
	fmt.Print("Enter the number of days: ")
	fmt.Scanln(&days) //modify the value of days to the user input

	fmt.Printf("The fine is: %.2f\n", libFines(days))
}

// takes a number of days
// return a float value of fines
func libFines(days int) float32 {
	var fine float32
	const finePerDay float32 = 0.1
	const finePerDayAfter7 float32 = 0.2

	if days <= 7 {
		fine = finePerDay * float32(days)
	} else {
		fine = finePerDayAfter7 * float32(days)
	}

	return fine
}
