package main

import "strconv"

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
