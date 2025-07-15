package main

import (
	"fmt"
)

func MergeAlternately(word1 string, word2 string) string {
    var result = make([]rune, 0, len(word1) + len(word2))
	oneRunes := []rune(word1)
	twoRunes := []rune(word2)

	maxLen := len(oneRunes) + len(twoRunes)

	oneLen := len(oneRunes)
	twoLen := len(oneRunes)

	// for i := 0; i < maxLen; i++ {
	// 	fmt.Println(i % 2)
	// }

	k := 0 
	n := 0
	
    for i := 0; i < maxLen; i++ {
		fmt.Println(k)
		if i % 2 == 0 && i < oneLen {
			result = append(result, oneRunes[k])
			k++
		}
		if i % 2 == 1 && i < twoLen {
			result = append(result, twoRunes[n])
			n++
		}
	}
	return string(result)
}

func main() {
	var inpStr1 string
	var inpStr2 string

	fmt.Scan(&inpStr1, &inpStr2)

	result := MergeAlternately(inpStr1, inpStr2)
	fmt.Print("\n")
	fmt.Print(result)
}
