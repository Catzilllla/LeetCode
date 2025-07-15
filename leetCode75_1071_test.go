package main

import (
	"testing"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
    var x string
    lenStr1 := len(str1)
    lenStr2 := len(str2)

	/* чекер строки1 с множеством повторений */
	sameChar := str1[0]
	for _, char := range str1 {
		if 
	}

	/* чекаем что первый символ != первому символу 2ой лексемыы */
    if lenStr1 > lenStr2 {
        if str2[0] != str1[lenStr2] {
            return ""
        }
        x = str1[lenStr2:]
    }

    if lenStr1 < lenStr2 {
		s := strings.
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

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		name string
		word1 string
		word2 string
		result string
	}{
		{
			name:	"Test 1",
			word1:	"LEET",
			word2:	"CODE",
			result: "",
		},
		{
			name:	"Test 2",
			word1:	"ABAB",
			word2:	"AB",
			result: "AB",
		},
		{
			name:	"Test 3",
			word1:	"ABCABC",
			word2:	"ABC",
			result: "ABC",
		},
		{
			name:	"Test 4",
			word1:	"AAAAAAAA",
			word2:	"A",
			result: "A",
		},
		{
			name:	"Test 5",
			word1:	"TAUXXTAUXXTAUXXTAUXXTAUXX",
			word2:	"TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX",
			result: "TAUXXTAUXXTAUXXTAUXXTAUXX",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcdOfStrings(tt.word1, tt.word2)


			// 1 <= str1.length, str2.length <= 1000
			// строки str1 и str2 состоят из английских прописных букв.

			if len(tt.word1) < 1 || len(tt.word2) > 1000 {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.result)
				return
			}
			for i := range got {
				if got[i] != tt.result[i] {
					t.Errorf("gcdOfStrings() = %v, want %v", got, tt.result)
				}
			}
		})
	}
}