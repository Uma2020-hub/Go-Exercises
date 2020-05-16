package main

import (
	"fmt"
	"math"
)

const yearVal, weekVal int = 365, 7

//CalculateDays - calculate days of week and days of year
func CalculateDays(days int) (int, int, int) {
	var daysinWeek, daysinYear int
	if days > 0 {
		daysinYear = days / yearVal
		//week = (days-(year * 365))/7
		daysinWeek = (days - (daysinYear * yearVal)) / weekVal
		//daysinWeek = days / weekVal

	}
	return days, daysinWeek, daysinYear
}

/* func ConvertDegreeToCelsius(celsius int)int{

} */

//CalculateSqrt - Calculate SquareRoot of a number
func CalculateSqrt(sqrtVal float64) float64 {
	return math.Sqrt(float64(sqrtVal))

}

//CelsiusToFarenHeit - Calculate degree celsius to Farenheit
func celsiusToFarenHeit(degreeVal int) int {
	return (degreeVal * 9 / 5) + 32
}

func main() {
	var numberOfDays, degreeCelsius, farenheit int
	var squareNumber, sqrtValue float64

	fmt.Println("Enter the number of days")
	fmt.Scanln(&numberOfDays)
	days, weeks, year := CalculateDays(numberOfDays)
	fmt.Println("Enter the number to find a square root")
	fmt.Scanln(&squareNumber)
	if squareNumber > 0 {
		sqrtValue = CalculateSqrt(squareNumber)
	}
	fmt.Println("Enter the number in degree celsius")
	fmt.Scanln(&degreeCelsius)
	if degreeCelsius > 0 {
		farenheit = celsiusToFarenHeit(degreeCelsius)
	}
	fmt.Printf("\n-----------------------------------------------------------------------------------")
	fmt.Printf("Days Represented in \n Days %d \n Week %d \n Years %d", days, weeks, year)
	fmt.Printf("\n**********************************************************************************")
	fmt.Printf("\nSquare Root of the number %f is %f", squareNumber, sqrtValue)
	fmt.Printf("\n###################################################################################")
	fmt.Printf("\nFarenheit of  degree celsius %d is %d", degreeCelsius, farenheit)
	fmt.Printf("\n-----------------------------------------------------------------------------------")
}
