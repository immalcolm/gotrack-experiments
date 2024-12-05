package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//get user prompt for dewpoint temp
	var dewpoint float64
	fmt.Print("Enter the dewpoint temperature in Celsius: ")

	//read user input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input) //remove leading and trailing spaces

	dewpoint, err = strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The dewpoint temperature is: ", dewpoint)

	//calculate pressure
	//math.Pow(base, exponent)
	vaporPressure := 6.11 * math.Pow(10, (7.5*dewpoint)/(237.3+dewpoint))
	fmt.Printf("The vapor pressure is %.2f millibars\n", vaporPressure)
}
