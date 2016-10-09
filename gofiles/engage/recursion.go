package main

import "fmt"

func main() {
	fmt.Println("here we discuss recursion in go")

	fmt.Println(fact(7))

}

// This fact function calls itself until it reaches the base case of fact(0).

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
