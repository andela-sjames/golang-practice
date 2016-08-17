package main

import "fmt"

func main() {
	var arraySize int
	var numSlice []int
	var arrayInt int

	fmt.Scanf("%v", &arraySize)

	for i := 0; i < arraySize; i++ {
		fmt.Scanf("%v", &arrayInt)
		numSlice = append(numSlice, arrayInt)
	}

	sum := 0
	for _, num := range numSlice {
		sum += num
	}

	fmt.Println(sum)

}
