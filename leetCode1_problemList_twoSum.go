package main

import (
	"testing"
)

/* сложность квадратичная -  переделать на map[] чтобы было O(n) */

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
