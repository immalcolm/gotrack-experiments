package main

import (
	"fmt"
	"sort"
)

func main() {
	//array
	var namesOfStudents [3]string
	//explicitly assign values
	namesOfStudents[0] = "John"

	fmt.Println(namesOfStudents[0])
	fmt.Println(namesOfStudents)
	fmt.Println(len(namesOfStudents))

	var numbers = [3]int{1, 2, 3}
	fmt.Println(numbers)

	fmt.Println("len of numbers: ", len(numbers))

	//slices
	var slicedNames = []string{"John", "Doe", "Jane"}
	fmt.Println(slicedNames)

	slicedNumbers := []int{1, 2, 3}
	fmt.Println(slicedNumbers)

	//using make
	var makeSlice = make([]int, 5)
	fmt.Println("Making slice", makeSlice)

	//loop
	for i := 0; i < len(slicedNames); i++ {
		fmt.Println(slicedNames[i])
	}

	for i, value := range slicedNames {
		fmt.Print("Sliced Names %d %s\n", i, value)
	}

	//append
	slicedNames = append(slicedNames, "Doe")
	fmt.Println(slicedNames)

	sort.Strings(slicedNames)
	sort.Ints(slicedNumbers)

	fmt.Println(slicedNames)
	fmt.Println(slicedNumbers)
}
