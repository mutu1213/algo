package other

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/valid-parentheses/submissions/
func isBracketValid(s string) bool {
	stack := make([]byte, 0)
	cp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			// 要考虑先进来右边括号的情况
			if len(stack) == 0 || stack[len(stack)-1] != cp[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}
	// 考虑最后只留下左边的情况
	if len(stack) == 0 {
		return true
	}
	return false
}

func Test_isBracketValid(t *testing.T) {
	fmt.Println(isBracketValid("{[]}"))
}
