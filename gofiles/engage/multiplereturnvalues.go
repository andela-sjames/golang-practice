//Go has built-in support for multiple return values. This feature is used often in idiomatic Go,
// for example to return both result and error values from a function.

package main

import "fmt"

func main() {
	fmt.Println("this is where we go further")

	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

	_, c := values()
	fmt.Println(c)

}

// The (int, int) in this function signature
// shows that the function returns 2 ints.

func values() (int, int) {
	return 7, 4
}
