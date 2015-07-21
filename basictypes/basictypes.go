package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// var declatations can be factored into blocks like import statements
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

/*
Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

// constant can be character, string, boolean or numeric values
// constant cannot be declared with short assignment
const Pi = 3.14

// like var statements, we can use block assignment here
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// %T is type, %v is value
	const template = "%T(%v)\n"
	fmt.Printf(template, ToBe, ToBe)
	fmt.Printf(template, MaxInt, MaxInt)
	fmt.Printf(template, z, z)

	// variable declared without initial values are given zero values
	// 0 for numeric values
	// false for boolean type
	// "" for strings
	var i1 int
	var f1 float64
	var b1 bool
	var s1 string
	fmt.Printf("%v %v %v %q\n", i1, f1, b1, s1)

	// Go assignments require explicit conversion when types are different
	var i2 int = 42
	var f2 float64 = float64(i2)
	var u2 uint = uint(f2)
	fmt.Println(i2, f2, u2)
	// more simply,
	// i2 := 42
	// f2 := float64(ii)
	// u2 := uint(ff)

	var x3, y3 int = 3, 4
	var f3 float64 = math.Sqrt(float64(x3*x3 + y3*y3))
	var z3 int = int(f3)
	fmt.Println(x3, y3, z3)

	// when no types are specified, i.e. not present in var declaration or in short assignment
	// and if the right-hand side variable is typed, then the type will be used
	var i4 int
	j4 := i4 // int

	// if the right-hand side variable is not typed
	// then the variable maybe int, float64 or complex128 depending on the precision of the constant
	i5 := 42
	f5 := 3.142
	g5 := 0.687 + 0.5i

	fmt.Printf("%T %T %T %T", j4, i5, f5, g5)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// numeric constants are high-precision `values`
	// an un-typed constant takes the type needed by its context
	fmt.Println(needInt(Small))   // convert const value to int type
	fmt.Println(needFloat(Small)) // convert const value to float64 value
	// fmt.Println(needInt(Big)) // constant 1267650600228229401496703205376 overflows int
	fmt.Println(needFloat(Big)) // convert const value to float64 value
}
