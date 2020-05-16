package main

import "fmt"

func main() {
	cmlStr := "oneTwoThree"
	var wordCount int
	var byteArray []byte
	byteArray = []byte(cmlStr)
	//var wordArray []string
	fmt.Println(byteArray)
	j := 0
	k := 0
	for i, v := range byteArray {
		//	wordArray = make([]string, len(byteArray))
		if v >= 65 && v < 97 {
			//	fmt.Println("k value ", k, i, v)
			if i == 0 && k == 0 {
				continue
			} else {
				fmt.Println(string(byteArray[k:i]))
			}

			//fmt.Println(i, v)

			//if i > 0 {

			//test := string(byteArr,ay[k:i])
			//wordArray[j] = cmlStr[:i]
			//wordArray = append(wordArray, cmlStr[:i])
			//	wordArray[j] = ([]byte)(byteArray[:i])
			//wordArray[j] = byteArray[:i]
			j++
			k = i
			//fmt.Println("test ", test)
			//	}

			//
			//wordArray[j] = string(byteArray[0:i])
			//j++
			wordCount++
		}
	}
	//fmt.Println("Test Word Array ", wordArray)
	//if k == len(byteArray) {
	wordCount++
	fmt.Println(string(byteArray[k:]))
	//}
	fmt.Println(wordCount)

}
