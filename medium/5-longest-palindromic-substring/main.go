package main

// https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	max := 0
	maxStr := ""
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			if checkPalindrome(s[i:j]) == true && j-i > max {
				max = j - i
				maxStr = s[i:j]
			}
		}
	}

	return maxStr
}

func checkPalindrome(s string) bool {
	j := len(s) - 1
	for i := 0; i <= j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
