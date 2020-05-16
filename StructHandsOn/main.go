package main

import (
	"fmt"
	"time"
)

//Car struct
type Car struct {
	model           string
	year            int
	miles           float64
	manufactureDate time.Time
	vin             string
	owner           string
}

//Employee struct
type Employee struct {
	name        string
	age         int
	designation string
}

func main() {

	var carModel Car

	carModel.owner = "Uma"
	carModel.model = "test"
	carModel.year = 2009
	carModel.miles = 12.5
	carModel.manufactureDate = time.Now()
	carModel.vin = "test"

	var emp Employee
	emp.age = 25
	emp.designation = "Technical Lead"
	emp.name = "Harini"

	printEmpValues(emp)
	printCarValues(carModel)
	modifyEmpValues(&emp)
	modifyCarValues(&carModel)
}

func printCarValues(element Car) {

	fmt.Printf("\n%#v \n", element)

}

func printEmpValues(element Employee) {

	fmt.Printf("\n%+v \n", element)

}

func modifyCarValues(element *Car) {
	element.owner = "Selva"
	fmt.Printf("\n%#v \n", *element)
}

func modifyEmpValues(element *Employee) {
	element.name = "Oviya"
	fmt.Printf("\n%#v \n", *element)
}
