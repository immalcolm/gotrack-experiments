package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now()
	fmt.Println("The time is", current_time)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at", t)
	fmt.Println(time.ANSIC)

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Feb  3 04:05:06 2015")

	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
	fmt.Printf("The value of parsedTime is %s\n", parsedTime)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())
}
