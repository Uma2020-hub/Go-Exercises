package main

import "fmt"

const (
	a = 1 << (10 * iota)
	b
	c
)
const (
	a2 = 1 << iota
	b2
	c2
	d2
)

func main() {
	var roles byte = a2 | c2 | d2
	fmt.Printf("%v \n", a2)
	fmt.Printf("%v \n", b2)
	fmt.Printf("%v  %v\n", c2, d2)
	fmt.Printf("%b \n", roles)
	fmt.Printf("%v is a1?", a2&roles == a2)
	fmt.Printf("%v is a1?", b2&roles == b2)

}
