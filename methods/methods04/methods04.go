package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// pointer receiver, therefore
// *Vertex implement Abser, not Vertex
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// an interface type is defined by a set of methods
type Abser interface {
	Abs() float64
}

func main() {
	f := MyFloat(-math.Sqrt2) // MyFloat
	v := Vertex{3, 4}         // Vertex

	var a, b Abser // interface
	/*
	   a value of interface type can hold any value of type that implements those methods
	   (analogy to OOP: type -> class, value -> instance)
	   it's the same as:
	       a := MyFloat(-math.Sqrt2) // MyFloat
	       a := Vertex{3, 4}         // Vertex
	*/
	a = f  // MyFloat implements Abser
	b = &v // *Vertex implements Abser, get pointer to use pointer receiver

	// this is an error because Vertex (not *Vertex) does not implement Abser
	// b = v
	fmt.Println(a.Abs())
	fmt.Println(b.Abs())
}
