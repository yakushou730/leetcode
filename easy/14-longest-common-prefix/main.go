package main

// https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	var prefixStr string
	var currMinLen int
	for i, str := range strs {
		if i == 0 {
			prefixStr = str
			currMinLen = len(prefixStr)
		} else {
			if len(str) == 0 {
				return ""
			}
			if len(prefixStr) > len(str) {
				currMinLen = len(str)
			}
			for j := 0; j < currMinLen; j++ {
				if prefixStr[j] == str[j] {
					if len(str) == currMinLen && len(str) == j+1 {
						prefixStr = str[0 : j+1]
						break
					}
					continue
				} else {
					prefixStr = str[0:j]
					currMinLen = len(prefixStr)
					break
				}
			}
		}
	}
	return prefixStr
}
