package main

import "fmt"

func main() {

	var num, fiboSeries int = 0, 0
	temp1 := 0
	temp2 := 1
	fmt.Println("Enter the number ")
	fmt.Scanln(&num)
	fmt.Println("Fibonacci Series")
	for i := 1; i <= num; i++ {
		if i == 1 {
			fmt.Print(" ", temp1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", temp2)
			continue
		}
		//temp := 0

		fiboSeries = temp1 + temp2
		temp1 = temp2
		temp2 = fiboSeries
		fmt.Printf("%d ", fiboSeries)

	}

}
