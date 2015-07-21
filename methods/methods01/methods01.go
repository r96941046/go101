package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
   Go does not have classes
   we define methods on structs (in fact, we can define methods on all types, all these types are like classes in OOP)
   the method receiver, which is like instances for class in OOP, `v *Vertex`, appears before method name and in parenthesis
*/

// pointer receiver here
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// v is a pointer to Vertex
	v := &Vertex{3, 4}
	// v := Vertex{3, 4}, result is the same?
	fmt.Println(v.Abs())
}
