package main

import (
	"testing"
)

func KidsWithCandies(candies []int, extraCandies int) []bool {
	// For each kid check if candies[i] + extraCandies ≥ maximum in Candies[i].
	var result = make([]bool, 0, 0)

	type Tbuffer struct {
		buffer []int
		maxCand int
	}

	newBuffer := Tbuffer{
		buffer: make([]int, 0, 0),
		maxCand: 0,
	}

	for _, cand := range candies {
		newBuffer.buffer = append(newBuffer.buffer, cand + extraCandies)
	}

	/* check max candies in candies[]int */
	for _, buff := range candies {
		if newBuffer.maxCand < buff {
			newBuffer.maxCand = buff
		}
	}

	for _, value := range newBuffer.buffer {
		if value >= newBuffer.maxCand {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		name string
		candiesKids []int
		extraCandies int
		result []bool
	}{
		{
			name:	"Test 1",
			candiesKids: []int{2,3,5,1,3},
			extraCandies: 3,
			result: []bool{true,true,true,false,true},
		},
		{
			name:	"Test 2",
			candiesKids: []int{4,2,1,1,2},
			extraCandies: 1,
			result: []bool{true,false,false,false,false},
		},
		{
			name:	"Test 3",
			candiesKids: []int{12, 1, 12},
			extraCandies: 10,
			result: []bool{true,false,true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KidsWithCandies(tt.candiesKids, tt.extraCandies)
			
			// n == candies.length
			if len(tt.candiesKids) != len(tt.result) {
				t.Errorf("len(tt.candiesKids) = %v and len(tt.result) = %v", len(tt.candiesKids), len(tt.result))
				return
			}
			// 2 <= n <= 100
			if len(tt.candiesKids) > 100 || len(tt.candiesKids) < 2 {
				t.Errorf("len(tt.candiesKids) = %v", len(tt.candiesKids))
				return
			}
			// 1 <= candies[i] <= 100
			for i := range got {
				if tt.candiesKids[i] > 100 || tt.candiesKids[i] < 1 {
					t.Errorf("tt.candiesKids[i] = %v", tt.candiesKids[i])
				}
			}
			// 1 <= extraCandies <= 50
			if tt.extraCandies < 1 || tt.extraCandies > 50 {
				t.Errorf("tt.extraCandies = %v", tt.extraCandies)
				return
			}

			for i := range got {
				if got[i] != tt.result[i] {
					t.Errorf("KidsWithCandies() = %v, want %v", got, tt.result)
				}
			}
		})
	}
}