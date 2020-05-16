package main

import "fmt"

func main() {
	var num1, num2 int
	var operator string
	fmt.Println("Enter the operator of your choice to do Arithmetic Operation (+,-,*,/)")
	fmt.Scan(&operator)

	fmt.Println("\nEnter the whole number")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Printf("Result is %d ", addArithmetic(num1, num2))
	case "-":
		fmt.Printf("Result is %d ", subtractArithmetic(num1, num2))
	case "*":
		fmt.Printf("Result is %d ", multiplyArithmetic(num1, num2))
	case "/":
		fmt.Printf("Result is %d ", divideArithmetic(num1, num2))
	}

}
func addArithmetic(i int, j int) int {
	return i + j
}

func subtractArithmetic(i int, j int) int {
	return i - j
}
func multiplyArithmetic(i int, j int) int {
	return i * j
}
func divideArithmetic(i int, j int) int {
	return i / j
}
