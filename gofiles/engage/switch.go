package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("here we discuss switch statement in go")

	i := 2
	fmt.Println("write", i, "as ")

	fmt.Println(time.Now().Weekday())

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's way pass noon")
	}
}
