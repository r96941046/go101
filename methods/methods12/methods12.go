package main

import (
	"io"
	"os"
	"strings"
)

/*
   a common pattern: an io.Reader wraps another io.Reader

   rot13Reader is a type that implements io.Reader
   it wraps another io.Reader, and customizes Read() method from it (like inheritance in OOP)
*/

type rot13Reader struct {
	r io.Reader
}

// implements Reader interface
func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	// uses the Reader interface r in its own struct
	n, err = rot.r.Read(p)
	// customize Read() method with rot13
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >= 'a' && p[i] < 'n') {
			p[i] += 13
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	return
}

func main() {
	// r uses the reader interface in its struct, which has a Read() method,
	// to implement and customize its own Read() method

	// s is an interface type
	s := strings.NewReader("Lbh penpxrq gur pbqr!")

	// pass the interface to struct r
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // You cracked the code!
}
