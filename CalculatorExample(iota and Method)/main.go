package main

import (
	"fmt"
	"math"
)

const (
	opAdd = iota
	opSub
	opMul
	opDiv
	opSqrt
)

func (c *calculator) Do(option int, op2 float64) {
	switch option {
	case opAdd:
		c.acc = c.acc + op2
		fmt.Printf("Add Value = %f", c.acc)

	case opSub:
		c.acc = c.acc - op2
		fmt.Printf("Sub Value = %f", c.acc)
	case opMul:
		c.acc = c.acc * op2
		fmt.Printf("Multiply Value = %f", c.acc)
	case opDiv:
		c.acc = c.acc / op2
		fmt.Printf("Division Value = %f", c.acc)
	case opSqrt:
		c.acc = math.Sqrt(c.acc)
		fmt.Printf("Sqrt Value = %f", c.acc)
	}
}

func main() {
	var option int
	var operand2 float64
	fmt.Println("Enter options 0 : Add 1 : Sub 2 : Mul 3: Div 4: Sqrt")
	fmt.Scanln(&option)
	fmt.Println("Enter the other operand")
	fmt.Scanln(&operand2)
	var cal = calculator{acc: 10}

	p := &cal
	//	cal.acc = 10
	p.Do(option, operand2)

}

type calculator struct {
	acc float64
}
