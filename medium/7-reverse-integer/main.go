package main

import "strconv"

// https://leetcode.com/problems/reverse-integer/
func reverse(x int) int {
	minus := 1
	if x < 0 {
		minus = -1
		x *= -1
	}

	str := strconv.Itoa(x)
	zeroCount := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == 0 {
			zeroCount += 1
		} else {
			break
		}
	}

	str = str[0 : len(str)-zeroCount]
	var resStr []uint8
	for i := len(str) - 1; i >= 0; i-- {
		resStr = append(resStr, str[i])
	}

	res, _ := strconv.Atoi(string(resStr))

	if (res > (1<<31)-1) || res < (-1*(1<<31)) {
		res = 0
	}

	return res * minus
}
