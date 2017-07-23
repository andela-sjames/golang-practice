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
// Go offers extensive support for times and durations;
// here are some examples.

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // We'll start by getting the current time.
    now := time.Now()
    p(now)

    // You can build a `time` struct by providing the
    // year, month, day, etc. Times are always associated
    // with a `Location`, i.e. time zone.
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // You can extract the various components of the time
    // value as expected.
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // The Monday-Sunday `Weekday` is also available.
    p(then.Weekday())

    // These methods compare two times, testing if the
    // first occurs before, after, or at the same time
    // as the second, respectively.
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // The `Sub` methods returns a `Duration` representing
    // the interval between two times.
    diff := now.Sub(then)
    p(diff)

    // We can compute the length of the duration in
    // various units.
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // You can use `Add` to advance a time by a given
    // duration, or with a `-` to move backwards by a
    // duration.
    p(then.Add(diff))
    p(then.Add(-diff))
}
