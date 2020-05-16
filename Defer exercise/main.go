package main

import "fmt"

func main() {
	var grades = [...]int{12, 13, 15, 17, 45}
	b := grades[1:4]
	b[2] = 34
	fmt.Printf("%v %v ", grades, b)
	fmt.Printf(" %v %v", len(grades), cap(grades))
	var a = make([]int, 3, 170)
	a = append(a, 89, 34, 66, 77)
	a = append(a, 99)

	fmt.Printf("%v %v ", a, b)
	fmt.Printf(" %v %v", len(a), cap(a))

}
