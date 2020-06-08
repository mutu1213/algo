package string

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

//https://leetcode-cn.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {
	var result int
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	neg := false
	if str[0] == '-' {
		neg = true
	}
	if !isNumber(str[0]) && str[0] != '-' && str[0] != '+' {
		return 0
	}
	start := 0
	if !isNumber(str[0]) {
		start = 1
	}
	for i := start; i < len(str); i++ {
		if isNumber(str[i]) {
			result = result*10 + int(str[i]-'0')
			if !neg && result > math.MaxInt32 {
				return math.MaxInt32
			}
			if neg && result*(-1) < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	if neg {
		result = result * (-1)
	}
	return result
}

func isNumber(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("9223372036854775808"))
}
