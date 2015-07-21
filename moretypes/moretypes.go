package main

import (
	"fmt"
	"math"
)

// a `struct` is a collection of fields
type Vertex struct {
	X int
	Y int
}

func demo_pointers() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p) // 42
	*p = 21
	fmt.Println(i) // 21

	p = &j
	*p = *p / 37
	fmt.Println(j) // 73
}

type Point struct {
	X, Y int
}

type Position struct {
	Lat, Long float64
}

// various struct literals
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0, Y:0
	p3 = &Vertex{1, 2} // has type *Vertex
)

func printSlice(s string, x []int) {
	// a slice is a reference to a part of an array
	// a slice's capacity represents the size of that backing array
	// so the length of a slice cannot be greater than the capacity
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

// closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// closure demo: fibonacci
func fibonacci() func() int {
	int1 := 0
	int2 := 1
	return func() int {
		result := int1
		int1, int2 = int2, int1+int2
		return result
	}
}

func main() {

	// the type `*T` is a pointer to a T-typed value
	// a pointer has `nil` zero value
	var p *int
	fmt.Println(p) // nil

	// & generates a pointer to its operand
	i := 42
	p = &i
	fmt.Println(p) // 0x208178178

	// * is the pointer's underlying value
	fmt.Println(*p) // 42
	*p = 21
	fmt.Println(i)  // 21
	fmt.Println(*p) // 21

	demo_pointers()

	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	// struct fields are accessed by using dots
	v.X = 4
	fmt.Println(v.X)

	// pointer to struct
	p2 := &v
	p2.X = 1e9 // not *p.X
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)

	// arrays
	// the type [n]T is an array of n values of type T
	// the length of an array is part of its type, so arrays cannot be resized
	var a1 [10]int   // array of ten integers
	var a2 [2]string // array of two strings
	a2[0] = "Hello"
	a2[1] = "World"
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a2[0], a2[1])

	// slices
	// []T is a slice with elements of type T
	// a slice points to an array of values and also includes a length
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	// slicing slices
	// slices can be re-sliced, creating a new slice value that points to the same array
	fmt.Println("s[1:4] ==", s[1:4]) // [3, 5, 7]
	fmt.Println("s[:3]", s[:3])      // [2, 3, 5]
	fmt.Println("s[4:]", s[4:])      //  [11, 13]

	// making slices with the make function
	// it allocates a zeroed array and then returns a slice that refer to that array
	a4 := make([]int, 5)
	printSlice("a4", a4)
	b4 := make([]int, 0, 5)
	printSlice("b4", b4)
	c4 := b4[:2]
	printSlice("c4", c4)
	d4 := c4[2:5]
	printSlice("d4", d4)
	e4 := a4[2:]
	printSlice("e4", e4)

	// nil slices
	// like pointers
	var z []int
	printSlice("z", z) // z len=0 cap=0 []
	if z == nil {
		fmt.Println("nil!")
	}

	// add elements to a slice
	var a5 []int
	printSlice("a5", a5)

	a5 = append(a5, 0)
	printSlice("a5", a5)

	// adding more than one element at once
	// if the backing array is too small, a bigger array will be allocated
	// the returned slice will point to the newly allocated array
	a5 = append(a5, 2, 3, 4)
	printSlice("a5", a5)

	// looping through an array
	// the range form of the for loop iterates over a slice or map
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("idx=%d 2**%d = %d\n", i, i, v)
	}

	// skipping for-range looping variables
	pow2 := make([]int, 10)
	// same as `for i := range pow2`
	for i, _ := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	// a map maps keys to values
	var m map[string]Position     // keys are string, values are Position
	m = make(map[string]Position) // maps must be created with `make` before use, the nil map is empty and cannot be assigned to
	m["Bell Labs"] = Position{
		40.68433, -74.39967,
	}
	fmt.Println(m)

	// map literals are like struct literals, but the keys are required
	var m2 = map[string]Position{
		"Bell Labs": Position{
			40.68433, -74.39967,
		},
		"Google": Position{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	// omitting type names
	var m3 = map[string]Position{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	// mutating maps
	m4 := make(map[string]int)
	// insert or update an element
	m4["Answer"] = 42
	fmt.Println(m4)
	m4["Answer"] = 48
	fmt.Println(m4)

	// delete an element
	delete(m4, "Answer")
	fmt.Println(m4)

	// get element with test
	// if key is present, ok is true; if not, ok is false, and v4 is the zero value of its type
	v4, ok := m4["Answer"]
	fmt.Println("The value", v4, "Present?", ok)

	// function references
	// fucntions are values too
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	// each closure remembers its own state
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	// fibonacci demo
	f5 := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f5())
	}
}
