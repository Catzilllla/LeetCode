package main

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
    buffer := nums[0]
    var bufferMass []int
    for i := 1; i < len(nums); i++ {
        if buffer + nums[i] == target {
            bufferMass = append(bufferMass, buffer)
            bufferMass = append(bufferMass, nums[i]) 
        }
    }
    return bufferMass
}

func testTwoSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		target int
		buffer []int
	}{
		{
			name:	"Test 1",
			nums:	[]int{2, 7, 11, 15},
			target:	9,
			buffer:	[]int{2, 7},
		},
		{
			name:	"Test 2",
			nums:	[]int{0, 0, 0, 0},
			target:	0,
			buffer:	[]int{0, 1},
		},
		{
			name:	"Test 3",
			nums:	[]int{1, 1, 1, 1},
			target:	2,
			buffer:	[]int{0, 1},
		},
		{
			name:	"Test 4",
			nums:	[]int{2, 7, 11, 15},
			target:	9,
			buffer:	[]int{2, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			if len(got) != len(tt.buffer) {
				t.Errorf("twoSun() = %v, want %v", got, tt.buffer)
				return
			}
			for i := range got {
				if got[i] != tt.buffer[i] {
					t.Errorf("twoSum() = %v, want %v", got, tt.buffer)
				}
			}
		})
	}
}