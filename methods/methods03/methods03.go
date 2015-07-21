package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

/*
methods can have a named type receiver or a pointer receiver
1. when it's a pointer receiver, i.e. func (p *Point):
    a. values are not copies on each method call
    b. the method can modify the value that the receiver points to
    c. this is more like an instance method in OOP, modifies instance directly
2. when it's a named type receiver, i.e. func (p Point):
    a. values are copied on each method call
    b. the method modifies the copied value but not the original value
    c. this is like an instance method in OOP that only processes copied values from the instance
*/

// a setter method differentiates between a named typed receiver and a pointer receiver
// when p is a Point, this Scale() method would not have any effect (values are copied)
func (p *Point) Scale(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

// a getter method is the same between a named typed receiver and a pointer receiver
func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	// p is a pointer to struct
	p := &Point{3, 4}
	p.Scale(5)
	fmt.Println(p, p.Abs())
}
