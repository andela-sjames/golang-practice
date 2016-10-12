// for is Go’s only looping construct. Here are three basic types of for loops.

package main

import "fmt"

func main() {
	fmt.Println("Here we discuss for loop in golang")

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly until you break
	// out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}
}
