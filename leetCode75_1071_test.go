package main

import (
	"testing"
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
			word1:	"ABCAB",
			word2:	"AB",
			result: "CAB",
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