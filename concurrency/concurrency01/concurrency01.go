package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < len(s); i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s, ":", i)
	}
}

func main() {
	/*
	   a goroutine is a lightweight thread managed by the Go runtime
	   go `f(x, y, z)` starts a new goroutine running `f(x, y, z)`

	   the evaluation of f, x, y, z happens in the current goroutine and the execution
	   of f happens in the new goroutine

	   goroutines run in the same address space, so access to shared memory must be synchronized
	   The sync package provides useful primitives, although you won't need them much in Go as there are other primitives
	*/

	// synchronous operation
	say("Hello")

	// asynchronous operation
	go say("world!")

	// asynchronous operation with an anonymous function call
	// this goroutine happens concurrently with the last one
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	/*
	   without this, the goroutines will not have time to run
	   because the program exits when the main() function finishes

	   one option would be to have the main goroutine block reading from a channel,
	   and have the goroutine write to the channel when it has completed its work.
	*/
	var input string
	fmt.Scanln(&input)
	fmt.Println("...done")
}
