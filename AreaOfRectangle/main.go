package main

import (
	"fmt"
)

func main() {
	var r rectangle
	p := &r
	p.setLengthAndWidth()
	area := p.getArea()
	perimeter := p.getPerimeter()
	fmt.Println("Area of Rectangle ", area)
	fmt.Println("Perimeter of Rectangle ", perimeter)
}

type rectangle struct {
	W, H float64
}

func (r *rectangle) setLengthAndWidth() {
	r.W = 10
	r.H = 4
}

func (r *rectangle) getArea() float64 {
	return r.W * r.H
}

func (r *rectangle) getPerimeter() float64 {
	//return 2 * pi * cr.Radius
	return 2*r.W + 2*r.H
}
