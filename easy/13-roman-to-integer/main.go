package main

import "fmt"

// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	romanMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var stringArr []string
	var tempStr string

	for _, ch := range s {
		str := string(rune(ch))
		if tempStr == "" {
			tempStr = str
			continue
		} else {
			if romanMap[tempStr] < romanMap[str] {
				stringArr = append(stringArr, fmt.Sprintf("%s%s", tempStr, str))
				tempStr = ""
			} else {
				stringArr = append(stringArr, tempStr)
				tempStr = str
			}
		}
	}

	if tempStr != "" {
		stringArr = append(stringArr, tempStr)
	}

	var sum int
	for _, str := range stringArr {
		sum += romanMap[str]
	}

	return sum
}
