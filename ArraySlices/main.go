package main

import "fmt"

func main() {
	var numberList [11]int = [11]int{2, 9, 3, 1, 4, 7, 5, 6, 7, 8, 15}
	var multiplesList []int
	for i := 0; i < 11; i++ {
		if numberList[i]%3 == 0 {
			multiplesList = append(multiplesList, numberList[i])
		}
	}
	sum := 0
	for _, v := range multiplesList {
		sum = sum + v
	}
	fmt.Printf("The sum of multiples of 3 is %d", sum)

}
