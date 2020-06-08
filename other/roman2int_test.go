package other

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/roman-to-integer/submissions/
//注意点，优先考虑两个字符代表一个数字的情况
func roman2Int(s string) int {
	var result int
	m := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100,
		"XC": 90, "L": 50, "XL": 40, "X": 10,
		"IX": 9, "V": 5, "IV": 4, "I": 1,
	}
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if v, ok := m[string(s[i])+string(s[i+1])]; ok {
				result = result + v
				i++
			} else {
				if v, ok := m[string(s[i])]; ok {
					result = result + v
				}
			}
		} else {
			if v, ok := m[string(s[i])]; ok {
				result = result + v
			}
		}

	}
	return result
}

func TestRoman2Int(t *testing.T) {
	fmt.Println(roman2Int("MCM"))
}
