package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number int
	fmt.Println("Enter any number")
	fmt.Scanln(&number)
	c1 := computerChoice(number)
	fmt.Println(c1)
	fmt.Println(getWinner(c1, number))

}
func computerChoice(num int) int {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r1.Intn(num)
}

func getWinner(num1 int, num2 int) string {
	computerRock := num1 < 33
	testRock := num2 < 33
	computerPaper := num1 < 66
	testPaper := num2 < 66
	computerScissor := num1 < 99
	testScissor := num2 < 99
	var result string
	if (computerRock && testRock) || (computerPaper && testPaper) || (computerScissor && testScissor) {
		result = "Match Draw"
	} else if (computerRock && testPaper) || (computerPaper && testScissor) || (computerScissor && testRock) {
		result = "Player Wins"
	} else {
		result = "Computer Wins"
	}
	return result

}
