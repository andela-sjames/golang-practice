//  Here we make attempt to iterate over values received from a channel.

package main

import "fmt"

func main() {
	fmt.Println("here we go again, let's iterate over a go channel")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
