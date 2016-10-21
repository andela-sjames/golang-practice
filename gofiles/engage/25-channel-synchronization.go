/*
We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.
*/

package main

import "fmt"
import "time"

func worker(done chan bool) {

	/*
	   This is the function we’ll run in a goroutine.
	   The done channel will be used to notify another
	   goroutine that this function’s work is done.
	*/

	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we’re done.
	done <- true
}

func main() {
	fmt.Println("Here we discuss go channel synchronization.")

	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done

	// If you removed the <- done line from this program,
	// the program would exit before the worker even started.
}
