package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

// make ErrNegativeSqrt an error by implementing Error()
func (e ErrNegativeSqrt) Error() string {
	// calling fmt.Sprint(e) here causes infinite loop
	// use fmt.Sprint(float64(e)) to avoid
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), nil
	}
	return 0, ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
