package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	// after calculation, send sum to channel c
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	/*
		channels are typed conduit through which you can send and receive values with the channel operator
		<-, ->, the data flows in the direction of the arrow
	*/

	// like maps and slices,
	// channels must be created before use
	c := make(chan int)

	// by default, sends and receives block until the other side is ready
	// this allows goroutines to synchronize without explicit locks or condition variables
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
