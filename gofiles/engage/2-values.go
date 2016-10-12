// Go has various value types including strings, integers, floats, booleans, etc.

package main

import "fmt"

func main() {
	// Strings, which can be added together with +.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// Booleans, with boolean operators as you’d expect.
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
}
