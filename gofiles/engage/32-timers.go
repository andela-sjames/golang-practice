package main

import (
	"fmt"
	"time"
)

// Timers represent a single event in the future.
// You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.

func main() {
	fmt.Println("here we go again, here we look at go timers")

	timer1 := time.NewTimer(time.Second * 2)

	//The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer expired.
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is that you can cancel the timer before it expires.
	// Here’s an example of that.

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
