package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Person implements `Stringer` interface from fmt package
/*
type Stringer interface {
    String() string
}
*/

// named type receiver here
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	// many packages including fmt look for the interface `Stringer` to print values
	// the String() method is used to print values

	// no need for pointer `&` here, String() uses named type receiver
	fmt.Println(a, z)
}
