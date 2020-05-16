package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number")
	fmt.Scanln(&num)
	for i := 2; i <= num; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime == true {
			fmt.Println(i)
		}
	}
}
