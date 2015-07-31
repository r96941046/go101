package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		// receive from tick channel
		case <-tick:
			fmt.Println("tick.")
		// receive from boom channel
		case <-boom:
			fmt.Println("BOOM!")
			return
		// the default case in a select is run if no other case is ready
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
