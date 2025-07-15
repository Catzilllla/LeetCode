package main

import (
	"testing"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
    return true
}

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name string
		word1 string
		word2 string
		result string
	}{
		{
			name:	"Test 1",
			word1:	"abc",
			word2:	"pqr",
			result: "apbqcr",
		},
		{
			name:	"Test 1",
			word1:	"abcd",
			word2:	"pq",
			result: "apbqcd",
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeAlternately(tt.word1, tt.word2)


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