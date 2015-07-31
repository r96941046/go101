package main

import (
	"fmt"
)

/*
	a sender can close a channel to indicate that no more values will be sent

	a receiver can test whether a channel is closed by assigning a second parameter to the expression
	v, ok := <- ch
	`ok` is `false` if there are no more values and the channel is closed
*/

// sender
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	// only the sender should close a channel, the receiver should never close a channel

	/*
		channels aren't like files, you don't usually need to close a channel
		it's only necessary when the receiver must be told there are no more values comming
		such as this case, terminating a range loop
		there will be an error is no close() is present in this example
	*/
	close(c)
}

// range and channel close
func main() {
	// a channel with buffer length 10
	c := make(chan int, 10)

	// gorouting with channel
	go fibonacci(cap(c), c)

	// for {
	// 	v, ok := <-c
	// 	fmt.Println(v, ok)
	// 	if ok == false {
	// 		break
	// 	}
	// }

	// receiver from channel
	// the loop receives values from the channel repeated until it is closed
	for i := range c {
		fmt.Println(i)
	}
}
