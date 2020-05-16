package main

import (
	"fmt"
	"math"
)

func main() {
	defer fmt.Println("In Main before calling function")
	fmt.Println(sqrt(-122))
	defer fmt.Println("In Main after calling function")
}

func sqrt(x float64) (float64, error) {
	var result float64
	var err error
	defer fmt.Println("the value of x is ", x)
	if x > 0 {
		result = math.Sqrt(x)
	} else {
		defer fmt.Println("x is a negative value")
		panic("Cannot enter negative numbers")
	}
	return result, err
}
