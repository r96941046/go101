package main

import (
	"fmt"
)

func main() {
	// this is a buffered channel, buffer length = 2
	// the buffer overkills when length = 1 in this example
	ch := make(chan int, 2)

	// sends to a buffered channel block only when the buffer is full
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
