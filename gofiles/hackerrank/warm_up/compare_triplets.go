package main

import "fmt"

func main() {
	a, b := ReadArrays()
	sa, sb := CompareIt(a, b)
	fmt.Printf("%d %d", sa, sb)
}

func CompareIt(a, b []int) (int, int) {
	var scoreA int
	var scoreB int

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			scoreA += 1
		} else if b[i] > a[i] {
			scoreB += 1
		}
	}
	return scoreA, scoreB
}

func ReadArrays() ([]int, []int) {
	a := make([]int, 3)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}
	b := make([]int, 3)
	for i := range b {
		fmt.Scanf("%d", &b[i])
	}
	return a, b
}
