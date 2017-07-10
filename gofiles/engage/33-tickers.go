// Timers are for when you want to do something once in the future -
// tickers are for when you want to do something repeatedly at regular intervals.
// Here’s an example of a ticker that ticks periodically until we stop it.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("here we go again as we define the go tickers.")

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
