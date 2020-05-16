package main

import "fmt"

// Filter uses given function to filter the values from given slice
func Filter(s []int, filter func(int) bool) []int {
	var result []int
	for _, v := range s {
		if filter(v) {
			result = append(result, v)
		}

	}
	return result
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 17, 29, 22, 24}
	filterOdd := func(x int) bool {
		// implement this to filter odd number
		if x%2 != 0 {
			return true
		}
		return false

	}
	s = Filter(s, filterOdd)
	fmt.Println(s)
	// filterPrime = func(int) bool {
	// implement this function to filter prime numbers

	// Filter(s, )
	//fmt.Println("Hello, playground")
}
