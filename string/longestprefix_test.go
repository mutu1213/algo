package string

import (
	"strings"
)

//https://leetcode-cn.com/problems/longest-common-prefix/submissions/
//思路：假设第一个字符串会起始最长子串，挨个比下去找最长子串，不匹配则减少prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			if prefix == "" {
				return prefix
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

//思路2：按列比，比到其中一个字符串为到头或者该列不全相同
func longestCommonPrefix1(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return ""
	}
	for j := 0; j < len(strs[0]); j++ {
		for i := 0; i < len(strs)-1; i++ {
			if j == len(strs[i]) || j == len(strs[i+1]) {
				return prefix
			}
			if strs[i][j] != strs[i+1][j] {
				return prefix
			}
		}
		prefix = strs[0][:j+1]
	}

	return prefix
}
