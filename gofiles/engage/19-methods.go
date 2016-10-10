// Go supports methods defined on struct types.

package main

import "fmt"

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect (pointer).
func (r *rect) area() int {
	return r.width * r.height
}

/* Methods can be defined for either pointer or value RECEIVER types.
Here’s an example of a value receiver. */
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println("here we discuss golang method...")

	r := rect{width: 10, height: 5}

	// call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	/*
	   Go automatically handles conversion between values and
	   pointers for method calls.You may want to use a pointer
	   receiver type to avoid copying on method calls or to allow
	   the method to mutate the receiving struct.
	*/

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
