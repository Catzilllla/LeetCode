package main

import (
	"testing"
	// "strings"
	"fmt"
)

func twoSum_1(nums []int, target int) []int {
    var bufferMass []int

	flag := 0
    for i := 0; i < len(nums); i++ {
        for k := 0; k < len(nums); k++ {
			// fmt.Println(nums[i], nums[k])
            if nums[i] + nums[k] == target && i != k {
                bufferMass = append(bufferMass, i)
                bufferMass = append(bufferMass, k)
				flag = 1
				break
            }
        }
		if flag == 1 {
			break
		}
    }
    return bufferMass   
}

func twoSum(nums []int, target int) []int {
	var buffer = make(map[int]int)

	for i, num := range nums {
		compliment := target - num
		if j, found := buffer[compliment]; found {
			return []int{j, i}
		}
		buffer[num] = i
		// fmt.Println(buffer, compliment)
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		target int
		output []int
	}{
		{
			name:	"Test 1",
			nums:	[]int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			name:	"Test 2",
			nums:	[]int{3,2,4},
			target: 6,
			output: []int{1, 2},
		},
		{
			name:	"Test 3",
			nums:	[]int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
		
			// 2 <= nums.length <= 104
			len := len(tt.nums)
			if len < 2 || len > 104 {
				t.Errorf("2 <= nums.length <= 104")
				return
			}
			// -109 <= nums[i] <= 109
			for i,_ := range tt.nums {
				if tt.nums[i] < -109 || tt.nums[i] > 109 {
					t.Errorf("nums: [ -109 <= nums[i] <= 109 ] :: nums[%d] = %v", i, tt.nums)
					return
				}
			}
			// -109 <= target <= 109
			if tt.target < -109 || tt.target > 109 {
				t.Errorf("-109 <= target <= 109")
				return
			}
			// main test
			for i,_ := range got {
				if got[i] != tt.output[i] {
					t.Errorf("twoSum() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}