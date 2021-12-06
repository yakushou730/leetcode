package main

import "strconv"

// https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3237/
func findNumbers(nums []int) int {
	sum := 0
	for _, num := range nums {
		str := strconv.Itoa(num)
		if len(str)%2 == 0 {
			sum += 1
		}
	}

	return sum
}
