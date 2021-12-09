package main

import "fmt"

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println(nums)
	return k
}
