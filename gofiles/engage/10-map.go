// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

package main

import "fmt"

func main() {
	fmt.Println("We have not even started yet")

	// To create an empty map, use the builtin make: make(map[key-type]val-type).

	m := make(map[string]int) // equivalent of python dict

	m["val1"] = 45
	m["val2"] = 90

	fmt.Println("map: ", m)

	v1 := m["val1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	// The builtin delete removes key/value pairs from a map.

	delete(m, "val2")
	fmt.Println("map: ", m)

	// The optional second return value when getting a value from a map indicates if
	// the key was present in the map. This can be used to disambiguate between missing
	// keys and keys with zero values like 0 or "". Here we didn’t need the value itself,
	// so we ignored it with the blank identifier _.

	_, prs := m["val2"]
	fmt.Println("prs: ", prs)

	// You can also declare and initialize a new map in the same line with this syntax.

	n := map[string]int{"foo": 1, "bar": 2} // create the map on the fly!
	fmt.Println("map: ", n)
}
