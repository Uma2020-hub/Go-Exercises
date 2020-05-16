package calculatepackage

import (
	"math"
)

const yearVal, weekVal int = 365, 7

//CalculateYearMonthAndDays - calculate days of week and days of year
func CalculateYearMonthAndDays(days int) (int, int, int) {
	var daysinWeek, daysinYear, remainingDays int
	if days > 0 {
		daysinYear = days / yearVal
		daysinWeek = (days - (daysinYear * yearVal)) / weekVal
		remainingDays = (days - (daysinYear * yearVal)) % weekVal
	}
	return remainingDays, daysinWeek, daysinYear
}

//CalculateSqrt - Calculate SquareRoot of a number
func CalculateSqrt(sqrtVal float64) float64 {
	return math.Sqrt(sqrtVal)

}
