package main

import (
	"fmt"
	"math"
)

func appendInt(a1 []int, v ...int) {
	i := 5
	for _, b := range v {
		a1[i] = b
		i++
	}
}

func main() {
	//b := make([][]int, 5)
	//a := make([]int, 3)
	//	a := []int{1, 2, 3}
	//b := [][]int{
	//	[]int{2, 3, 4},
	//	[]int{4, 5, 6},
	// }
	a := make([]int, 5, 10) // []int{1, 2, 4, 8, 16, 6, 7, 8}
	for i := range a {
		a[i] = int(math.Pow(2, float64(i)))
	}
	fmt.Println(len(a), cap(a))

	a = a[:10] // len 10, cap 10
	fmt.Println(len(a), cap(a))
	var k [7]int = [7]int{25, 26, 27, 28, 29, 30, 31}
	z := k[1:5]
	appendInt(a, 6, 7, 8, 9, 10)
	fmt.Println("Before K")
	appendInt(a, z...)
	fmt.Println(a)
}
