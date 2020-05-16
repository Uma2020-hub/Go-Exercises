package main

import "fmt"

func main() {
	var phy, che, bio, macs, comp int
	var markPercentage float32
	fmt.Println("Enter the physics, chemistry,biology,mathematics and computer")
	fmt.Scan(&phy, &che, &bio, &macs, &comp)
	markPercentage = (float32(phy+che+bio+macs+comp) / 500.0) * 100
	fmt.Printf("Total percentage %.2f \n", markPercentage)
	fmt.Println(" ")
	if markPercentage >= 90 {
		fmt.Println("Grade A")
	} else if markPercentage >= 80 {
		fmt.Println("Grade B")
	} else if markPercentage >= 70 {
		fmt.Println("Grade C")
	} else if markPercentage >= 60 {
		fmt.Println("Grade D")
	} else if markPercentage >= 40 {
		fmt.Println("Grade E")
	} else {
		fmt.Println("Grade F")
	}

}
