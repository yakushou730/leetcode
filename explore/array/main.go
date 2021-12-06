package main

func findMaxConsecutiveOnes(nums []int) int {
	counter := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter += 1
			if counter > max {
				max = counter
			}
		} else {
			counter = 0
		}
	}
	return max
}
