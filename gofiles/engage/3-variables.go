/*
In Go, variables are explicitly declared and used by
the compiler to e.g. check type-correctness of function calls.
*/

package main

import "fmt"

func main() {
	fmt.Println("here we go again")

	// var declares 1 or more variables.
	var a string = "This is Andela"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	// 0 for ints, 0.0 for floats, "" for strings, nil for pointers, …) We can also use the new function:
	// For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// For example, the zero value for a string is ' '.
	var letter string
	fmt.Println(letter)

	// The := syntax is shorthand for declaring and initializing a variable, e.g.
	// for var f string = "short" in this case.
	short := "This is for short declaration."
	fmt.Println(short)
}
