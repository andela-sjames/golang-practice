// Variadic functions can be called with any number of trailing arguments.

package main

import "fmt"

func main() {
	fmt.Println("here we see something different...")

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println("total: ", total)
}
