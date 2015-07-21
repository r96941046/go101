package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// implements built-in interface `error`
/*
type error interface {
    Error() string
}
*/
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didnt work",
	}
}

func main() {
	// check if there's error when executing functions: check if err is nil
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
