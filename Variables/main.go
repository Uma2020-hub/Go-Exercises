package main

import (
	"fmt"
	"math"
	"strconv"
)

const stringValue string = "756"
const yearVal, weekVal int = 365, 7

//CalculateDays - calculate days of week and days of year
func CalculateDays(days int) (int, int, int) {
	var daysinMonth, daysinYear int
	if days > 0 {
		daysinYear = days / yearVal
		daysinMonth = days / weekVal
	}
	return days, daysinMonth, daysinYear
}

//func Convert
func main() {
	var j, numberOfDays int
	var sqrtValue float64

	//var int result
	fmt.Println("Enter the number to find a square root")
	fmt.Scanln(&j)
	if j > 0 {
		sqrtValue = math.Sqrt(float64(j))
	}
	fmt.Printf("The Square Root of number %d  is %.2f", j, sqrtValue)
	fmt.Printf("\n The String representation of number %d of type %T is %v", j, j, strconv.Itoa(j))
	k, err := strconv.Atoi(stringValue)
	if err != nil {
		fmt.Println("\n Following Error Occured", err)
	} else {
		fmt.Printf("\n The String representation of string %v of type %T is %d", stringValue, stringValue, k*j)
	}
	fmt.Println("Enter then number of days")
	fmt.Scanln(&numberOfDays)
	days, months, year := CalculateDays(numberOfDays)
	fmt.Printf("Days Represented in \n Days %d \n Months %d \n Years %d", days, months, year)
}
