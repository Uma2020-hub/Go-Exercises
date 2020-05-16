package main

import (
	"fmt"
	"math"
)

//Circle ....
type Circle struct {
	r float64
}

//Rectangle ...
type Rectangle struct {
	x, y float64
}

type geometry interface {
	area() float64
	perim() float64
}

func (r Rectangle) area() float64 {
	return r.x * r.y
}

func (r Rectangle) perim() float64 {
	return 2*r.x + 2*r.y
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.r
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

//Measure ...
func Measure(g geometry) {
	//var i geometry
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	var r = Rectangle{x: 12, y: 13}
	var c = Circle{r: 6}
	/* 	var i geometry
	   	i = &c
	   	fmt.Println("Area and Perimeter of Circle")
	   	fmt.Println(i.area())
	   	fmt.Println(i.perim())

	   	fmt.Println("Area and perimeter of Rectangle")
	   	i = &r
	   	fmt.Println(i.area())
	   	fmt.Println(i.perim()) */
	Measure(c)
	Measure(r)
}
