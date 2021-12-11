package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		existMap := make(map[uint8]bool)
		for j := i; j < len(s); j++ {
			if _, ok := existMap[s[j]]; ok == true {
				if len(existMap) > max {
					max = len(existMap)
				}
				break
			} else {
				existMap[s[j]] = true
				if (j == len(s)-1) && len(existMap) > max {
					max = len(existMap)
				}
			}
		}
	}

	return max
}
