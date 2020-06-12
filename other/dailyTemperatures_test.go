package other

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/daily-temperatures/submissions/
//新make一个result相当于都赋了默认值0，o(n2）暴力解法
func dailyTemperatures(T []int) []int {
	if len(T) == 0 {
		return nil
	}

	result := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for j := i + 1; j < len(T); j++ {
			if T[j] > T[i] {
				result[i] = j - i
				break
			}
		}
	}
	return result
}

//单调栈,从底到顶，依次减小，时间复杂度是o(2n)，进栈出栈各一次，而不是在出栈N，注意，从stack里取出来的是T的下标
func dailyTemperatures1(T []int) []int {
	if len(T) == 0 {
		return nil
	}
	result := make([]int, len(T))
	var stack []int
	for i := 0; i < len(T); i++ {
		for j := len(stack) - 1; j >= 0 && T[i] > T[stack[j]]; j-- {
			result[stack[j]] = i - stack[j]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}

func Test_dailyTemperatures(t *testing.T) {
	fmt.Println(dailyTemperatures1([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
