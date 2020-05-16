package main

import (
	"fmt"
	"unicode"
)

func main() {
	//var originalStr = "abcdefghijklmnopqrstuvwxyz"
	//var byteStr = []byte(originalStr)
	var result string
	var testStr = "middle-Outz"
	k := 2
	//originalStr.append()
	/* convStr := byteStr[k:]
	newArray := byteStr[0:k]
	newStr = string(convStr) + string(newArray)
	var newByteArray = []byte(newStr)
	fmt.Println(string(newByteArray)) */
	/* for _, v := range testStr {
		for i, s := range originalStr {
			if v == s {
				//if unicode.IsUpper()
				result = result + string(newStr[i])
				//fmt.Println()
			}
		}
	} */
	for _, v := range testStr {
		if unicode.IsUpper(v) == true {
			l := ((int(v)+k-65)%26 + 65)
			result = result + string(l)
		} else if unicode.IsLower(v) == true {
			l := ((int(v)+k-97)%26 + 97)
			result = result + string(l)
		} else {
			result = result + string(v)
		}
	}
	fmt.Println(result)
	/* fmt.Println(testStr)
	fmt.Println([]byte(testStr))
	fmt.Println(result)
	fmt.Println([]byte(result)) */

}
