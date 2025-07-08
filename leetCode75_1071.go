package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	// for i := 65; i < 123; i++ {
	// 	if i < 91 || i > 96 {
	// 		fmt.Print(string(i))
	// 	}
	// }ß

	var buffer string
	if str1[0] != str2[0] {
		buffer = ""
		return buffer
	}
	for i, _ := range str2 {
		// fmt.Printf("%d %v\n", i, string(r))
		
		buffer = str1[i+1:]
		if str1[i] != str2[i] {
			return buffer
		}
	}
	return buffer
}

func main() {
	var inpStr1 string
	var inpStr2 string

	fmt.Scan(&inpStr1, &inpStr2)

	result := gcdOfStrings(inpStr1, inpStr2)
	fmt.Print("\n")
	fmt.Print(result)
}
