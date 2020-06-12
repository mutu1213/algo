package other

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/valid-parentheses/submissions/
func isBracketValid(s string) bool {
	m := make(map[byte]int)
	cp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			m[s[i]] = m[s[i]] + 1
		} else {
			if m[cp[s[i]]] == 0 {
				return false
			} else {
				m[cp[s[i]]] = m[cp[s[i]]] - 1
			}
		}
	}
	if m['('] == 0 && m['{'] == 0 && m['['] == 0 {
		return true
	}
	return false
}

func Test_isBracketValid(t *testing.T) {
	for i := 5; i < 13; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(fmt.Sprintf("ALTER table  tb_device_self_checks_2020%02d%03d  modify column msg text;", i, j))
		}
	}
	//fmt.Println(isBracketValid("()"))
}
