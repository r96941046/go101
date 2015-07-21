package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	// the if statement
	// no ()
	// {} is required
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, limit float64) float64 {
	// if statement can take a short statement to execute before the condition
	// variable declared by the statement is only visible in the if/else block
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit // v can not be used here
}

func newton_sqrt(x float64) float64 {
	// or guess: = float64(1)
	guess := 1.0
	for i := 0; i < 10; i++ {
		guess = (guess + (x / guess)) / 2.0
	}
	return guess
}

func defer_counter() {
	fmt.Println("Start counting...")

	for i := 0; i < 10; i++ {
		// deferred function calls are pushed onto a stack
		// they are executed in last-in-first-out order
		defer fmt.Println(i)
	}

	// this appears before the counting
	fmt.Println("Done counting.")
}

func main() {

	// defer statements defers the execution of a function until the surrounding function returns
	// defer's arguments are immediately evaluated, but the execution is deferred
	defer fmt.Println("End of execution of main")

	fmt.Println("Start of execution of main")

	sum := 0

	// Go has only one looping construct, the for loop
	// no ()
	// {} is required
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// pre and post statements can be ignored
	// and this is how `while` is spelled `for` in Go
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// when condition is omitted, it becomes infinite loop
	/*
		for {
		}
	*/

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20), // the comma here is required, or the semicolon will be inserted
	)

	fmt.Println(newton_sqrt(2))

	// a case body breaks automatically, unless it ends with a `fallthrough` statement
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	// switch evaluates case from top to bottom
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	// switch without condition is the same as `switch true`
	// so this is a good replacement for a long if/else chain
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer_counter()
}
