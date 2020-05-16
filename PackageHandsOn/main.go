package main

import (
	"calculatepackage"
	converter "calculatepackage/ConverterPackage"
	"fmt"
	//"fmt"
	//"fmt"
	//"converter"
)

func main() {
	fmt.Println(calculatepackage.CalculateYearMonthAndDays(100))
	fmt.Println(converter.CelsiusToFarenHeit(42))
}
