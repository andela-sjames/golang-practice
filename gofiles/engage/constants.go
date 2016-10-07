package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println("here we discuss constants in golang!")
	fmt.Println(s)

	const n = 500000
	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
