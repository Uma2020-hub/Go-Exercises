package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter any number")
	fmt.Scanln(&num)
	multipleOfNumber(num, 3)
	multipleOfNumber(num, 5)

}

func multipleOfNumber(n int, k int) {
	var total int
	fmt.Printf("\n Multiple of %d  : ", k)
	for j := 1; j <= n; j++ {
		fmt.Printf("%d ", k*j)
		total += k * j
	}
	fmt.Printf("Total of multiplication of %d  is %d", k, total)
}
