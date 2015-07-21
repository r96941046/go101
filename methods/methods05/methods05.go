package main

import (
	"fmt"
	"os"
)

/*
interfaces are satisfied implicitly
a type implements an interface by implementing the methods, no need to explicitly use `implements` keywords
this way, the implementation part is decoupled from the interface definition part,
and they don't need to be in the same package
*/

// package.io defines Reader and Writer, we don't have to
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	// os.Stdout implements Writer
	// i.e. value of os.Stdout has method Write
	var w Writer
	w = os.Stdout
	fmt.Fprint(w, "hello, writer\n")
}
