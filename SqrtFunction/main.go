package main

import (
	"fmt"
	"math"
)

func main() {
	sqrt(2)
	//fmt.Println(sqrt(2))
}

func sqrt(x float64) float64 {
	var z, result float64
	var temp float64
	if x < 0 {
		panic("Cannot enter negative numbers")
	} else {

		for z = x / 4; z <= 5.0; z++ {
			result = z - ((z*z - x) / (2 * z))
			fmt.Printf("\n Z square value of (z)%f is %f\n", z, result)
		}

		defer fmt.Printf("Actual sqrt  of (x) %f is %f", x, math.Sqrt(x))
		for z = x / 4; z <= 5.0; z++ {

			result = z - ((z*z - x) / (2 * z))
			if temp > 0 {
				variation := result - temp
				if variation < 1 {
					fmt.Printf("\n z value %f is close to x %f", z, x)
					fmt.Printf("\n Z square value of (z) %f is %f\n", z, result)
					break
				}
			}
			//	fmt.Printf("\n Z square value = %f for (z) is %.2f\n", result, z)
			temp = result
		}
	}
	return result
}
