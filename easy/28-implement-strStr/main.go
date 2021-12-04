package main

// https://leetcode.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	needleLen := len(needle)

	for i, ch := range haystack {
		if string(ch) == string(needle[0]) {
			if i+needleLen <= len(haystack) && haystack[i:i+needleLen] == needle {
				return i
			}
		}
	}
	return -1
}
