// Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a function inline without having to name it.

package main

import "fmt"

func main() {
	fmt.Println("here we go again bro! closures for real.")

	nextInt := intSeq()

	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value, which will be updated each time we call nextInt.

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function, create and test a new one.

	newInts := intSeq()
	fmt.Println(newInts())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
