package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	// select in an infinite loop
	for {
		/*
			select lets a gorouting wait on multiple communication operations
			1. a select blocks until one of its cases can run, and then it executes that case
			2. it chooses ont at random if multiple are ready

			select case must be receive, send of assign recv
		*/
		select {
		// case is send
		case c <- x:
			x, y = y, x+y
		// case is receive
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// goroutine with self-invoking function
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	// gorouting is still pending when the execution gets here
	fibonacci(c, quit)
}
