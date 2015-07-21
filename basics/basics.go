package main

import (
	"fmt"
	"github.com/r96941046/stringutil"
	"math"
	"math/rand"
	"time"
)

// at package level, outside functions, every statement should begin with a keyword
// like func, var, import, package...
// so the short assignment := is not available here

// the type comes after the variable declaration
func add(x int, y int) int {
	return x + y
}

// when two or more consecutive paratmeters share a type
// you can omit the type from all but the last
func minus(x, y int) int {
	return x - y
}

// a function can return any number of results, with added parenthesis
func swap(x, y string) (string, string) {
	return y, x
}

// return values could also be named for documentation purposes
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// naked return: return current values of the result, i.e. x, y
	// harms readability, and should be used only in short functions
	return
}

// list of variables, the type is last as well
var c, python, java bool

// initializers, when present, the type can be omitted
var j, k int = 1, 2

func main() {
	fmt.Println("Hello, World!")

	// By convention, package name is the same as the last element of the import path
	fmt.Println("The time is", time.Now())

	// The execution environment is deterministic
	// so rand.Intn always returns the same number
	// use rand.Seed to seed the number generator
	rand.Seed(11)
	fmt.Println("My favorite number is", rand.Intn(10))

	// A name is exported if it begins with a capital letter
	fmt.Println(math.Pi)

	// Custom function
	fmt.Println(add(42, 13))
	fmt.Println(minus(42, 13))

	// or just simply fmt.Println(swap("hello", "world"))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	// variable can be at package of function level
	var i int
	fmt.Println(i, c, python, java)

	// types are omitted here
	var ruby, javascript, says = true, false, "no!"
	fmt.Println(j, k, ruby, javascript, says)

	// short assignment can be used within function in place of var declarations
	// with implicit types
	s := 3
	l, m, n := true, false, "yes!"
	fmt.Println(s, l, m, n)

	fmt.Println("Now you have %g problems", math.Nextafter(2, 3))
	fmt.Printf(stringutil.Reverse("!oG ,olleH."))
}
