package main

// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	offset := 0
	var i int
	for {
		i = len(nums) / 2
		if i > 0 {
			if target > nums[i] {
				offset += i
				nums = nums[i:len(nums)]
			} else {
				nums = nums[0:i]
			}
			i = i / 2
		} else {
			if target > nums[i] {
				i += 1
			}
			break
		}
	}

	return offset + i
}
