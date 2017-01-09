package main

import "fmt"

func main() {
	fmt.Println("We have not even started yet")

	// make(map[key-type]val-type)

	m := make(map[string]int) // equivalent of python dict in a way

	m["val1"] = 45
	m["val2"] = 90

	fmt.Println("map: ", m)

	v1 := m["val1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "val2")
	fmt.Println("map: ", m)

	_, prs := m["val2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2} // create the map on the fly!
	fmt.Println("map: ", n)
}
