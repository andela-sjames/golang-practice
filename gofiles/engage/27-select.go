/*
Go’s select lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Here we discuss go select")

	// For our example we’ll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	/*
	   Each channel will receive a value after some amount of time, to simulate e.g.
	   blocking Remote Procedural Call(RPC) operations executing in concurrent goroutines.
	*/

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "One"

	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Two"
	}()

	/*
	   We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	*/

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
