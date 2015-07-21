package main

import (
	"fmt"
	"math"
)

// custom type (class in OOP)
type MyFloat float64

// we can define a method on any type that is declared in this package
// it's not allowed to define a method on a type from another package (including built in types)
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
