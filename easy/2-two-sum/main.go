package main

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	var firstIndex, secondIndex int
	ok := false
	for i := 0; i < len(nums); i++ {
		firstItem := nums[i]
		for j := i + 1; j < len(nums); j++ {
			secondItem := nums[j]
			if firstItem+secondItem != target {
				continue
			} else {
				ok = true
				firstIndex = i
				secondIndex = j
				break
			}
		}
		if ok {
			break
		}
	}

	return []int{firstIndex, secondIndex}
}
