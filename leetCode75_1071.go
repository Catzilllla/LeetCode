package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
    var x string
    lenStr1 := len(str1)
    lenStr2 := len(str2)
    // check := 0

    if lenStr1 > lenStr2 {
        if str2[0] != str1[lenStr2] {
            return ""
        }
        x = str1[lenStr2:]
    }

    if lenStr1 < lenStr2 {
        if str2[0] != str1[lenStr2] {
            return ""
        }
        x = str2[lenStr1:]
    }

	if lenStr1 == lenStr2 {
        if str1[0] != str2[0] {
            return ""
        }
    }

	return x
}

func main() {
	var inpStr1 string
	var inpStr2 string

	fmt.Scan(&inpStr1, &inpStr2)

	result := gcdOfStrings(inpStr1, inpStr2)
	fmt.Print("\n")
	fmt.Print(result)
}
