package main

import "fmt"

type animaler interface {
	Move(distance float64)
}
type stringer interface {
	print()
}
type dog struct {
	TotalDistance float64
}

func (d *dog) Move(dis float64) {
	d.TotalDistance = d.TotalDistance + dis
	//	d.print(d.TotalDistance)
	//fmt.Println(print(&d.TotalDistance))
}

func (d *dog) String() {
	fmt.Println("I moved  distance ", d.TotalDistance)
}

func main() {
	var d float64
	fmt.Println("Enter the distance")
	fmt.Scanln(&d)
	var p = dog{TotalDistance: 10}
	var i animaler
	i = &p
	i.Move(d)
	fmt.Println(d)
	//var s stringer
	//fmt.Println("test", s.print())
	//s = &p
	//s.print(11.30)
	//i.print(p.TotalDistance)
	//p.Move(d)

}
