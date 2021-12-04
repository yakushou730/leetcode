package main

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	var stack []int32
	for _, ch := range s {
		if len(stack) == 0 {
			stack = append(stack, ch)
		} else {
			lastItem := stack[len(stack)-1]
			switch ch {
			case '{', '[', '(':
				stack = append(stack, ch)
			case '}':
				if lastItem == '{' {
					stack = stack[0 : len(stack)-1]
				} else {
					stack = append(stack, ch)
				}
			case ']':
				if lastItem == '[' {
					stack = stack[0 : len(stack)-1]
				} else {
					stack = append(stack, ch)
				}
			case ')':
				if lastItem == '(' {
					stack = stack[0 : len(stack)-1]
				} else {
					stack = append(stack, ch)
				}
			}
		}
	}

	return len(stack) == 0
}
