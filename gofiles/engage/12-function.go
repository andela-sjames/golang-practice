package main

import "fmt"

func main() {
	fmt.Println("here we investigate golang functions.")
	var result int
	result = sum(3, 5)
	fmt.Println("result: ", result)

	sum := plusplus(3, 4, 5)
	fmt.Println("sum: ", sum)

	timesNum := makemultiple(4, 5)
	fmt.Println("multiplied value: ", timesNum)

}

func sum(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func makemultiple(a, b int) int {
	return a * b
}
