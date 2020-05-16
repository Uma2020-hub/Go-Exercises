package main

import (
	"fmt"
)

func main() {
	var ebUnit int
	var unitFees float32 = 0.0
	fmt.Println("Enter the number of Units")
	fmt.Scanln(&ebUnit)
	if ebUnit <= 50 {
		unitFees += float32(ebUnit) * .50
		//remainingUnit = ebUnit - 50
	}
	if ebUnit > 50 && ebUnit <= 150 {
		unitFees += ((float32(50) * .5) + float32(ebUnit-50)*.75)
	}
	if ebUnit > 150 && ebUnit <= 250 {
		unitFees += ((float32(50) * .5) + float32(100)*.75 + float32(ebUnit-150)*1.2)
	}
	if ebUnit > 250 {
		unitFees += ((float32(50) * .5) + float32(100)*.75 + float32(100)*1.2 + float32(ebUnit-250)*1.5)
	}

	unitFees += unitFees * 0.2

	fmt.Println("Total Current Bill", unitFees)

}
