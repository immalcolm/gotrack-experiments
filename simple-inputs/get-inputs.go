package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	//fmt.Fprintln(os.Stderr, err)
	if err != nil {
		fmt.Println(os.Stderr, err)
		//fmt.Println("You entered", aFloat)
	}
	fmt.Println("You entered", aFloat)
}
