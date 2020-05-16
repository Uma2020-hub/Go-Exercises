package main

import (
	"fmt"
	"time"
)

func main() {
	type car struct {
		model               string
		year                int
		miles               float64
		manufactureDateTime time.Time
		vin                 string
		owner               string
	}
	var c [4]car = [4]car{}
	var kidsCar = c[3:4]
	//var newCars = make([]car, 2, 4)

	c[0].miles = 14.5
	c[0].model = "Audi"
	c[0].vin = "vin"
	c[0].owner = "Uma"
	c[0].year = 1989
	c[0].manufactureDateTime = time.Now()

	c[1].manufactureDateTime = time.Now()
	c[1].miles = 198.5
	c[1].model = "BMV"
	c[1].vin = "vin"
	c[1].owner = "Harini"
	c[1].year = 1999

	c[2].manufactureDateTime = time.Now()
	c[2].miles = 124.5
	c[2].model = "Benz"
	c[2].vin = "vin"
	c[2].owner = "Selva"
	c[2].year = 1979

	c[3].manufactureDateTime = time.Now()
	c[3].miles = 0.0
	c[3].model = "Tesla"
	c[3].vin = "vin"
	c[3].owner = "Oviya"
	c[3].year = 2008
	fmt.Println("struct Array")
	for i := range c {
		fmt.Printf("\n%#v\n", c[i])
	}
	fmt.Println("kids Array")
	for i := range kidsCar {
		fmt.Printf("\n%#v\n", kidsCar[i])
	}
	/*	fmt.Println("struct Array")
		for i := range newCars {
			fmt.Printf("\n%#v\n", newCars[i])
		} */
}
