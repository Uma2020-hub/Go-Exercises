package main

import (
	"fmt"
)

func main() {
	var rightAns, wrongAns, totalAns = quizProgram()
	fmt.Printf("You have answered all %d Questions !", totalAns)
	fmt.Printf("\n you have answered %d correct answers and %d wrong answers ", rightAns, wrongAns)
	fmt.Printf("\n You got %d / %d", rightAns, totalAns)
}

func quizProgram() (int, int, int) {
	var userAns int
	var correctAnswers int
	var wrongAnswers int
	var totawlAnswers int
	var quizMap = map[string]int{
		"5+5": 10,
		"1+1": 2,
		"8+3": 11,
		"1+2": 3,
		"8+6": 14,
		"3+1": 4,
		"1+4": 5,
		"5+1": 6,
		"2+3": 5,
		"3+3": 6,
		"2+4": 6,
		"5+2": 7}

	for q, a := range quizMap {
		fmt.Printf("What is %v \n", q)
		fmt.Scanln(&userAns)
		if userAns == a {
			correctAnswers++
		} else {
			wrongAnswers++
		}
		totalAnswers = totalAnswers + 1
	}
	return correctAnswers, wrongAnswers, totalAnswers
}
