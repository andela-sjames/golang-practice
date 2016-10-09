package main

import "fmt"

func main() {
	fmt.Println(" here we discuss if-else statement in golang")

	// get ready to be zoomed.
	if 7%2 == 2 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4, yeah...")
	}

	if num := 9; num < 0 {
		fmt.Println("number is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println("number has multiple digits")
	}
}
