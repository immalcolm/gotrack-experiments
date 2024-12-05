package main

import "fmt"

func main() {
	fmt.Println("Morse Code Game")

	// Declare a variable of type string
	var userCode string

	//get user input for morse code
	fmt.Print("Enter a number to convert: ")

	//read user input
	fmt.Scanln(&userCode)

	convertedCode, count := morseCode(userCode)

	fmt.Println("The morse code is: ", convertedCode)
	fmt.Println("The number of digits converted is: ", count)

}

func morseCode(number string) (string, int) {

	//declare a map of morse code
	morseCodeMap := map[string]string{
		"0": "-----",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
	}

	var counter int
	var output string

	for _, v := range number {
		//check whether exists in the map
		_, exists := morseCodeMap[string(v)]
		if !exists {
			return "There contains an invalid input"
		} else {
			output += morseCodeMap[string(v)] + "   "
			counter++
		}
	}

	//convert user input to morse code
	return output, counter
}
