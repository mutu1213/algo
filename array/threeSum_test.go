package array

import (
	"sort"
	"testing"
)

//https://leetcode-cn.com/problems/3sum/submissions/
//思路：先数组排序，固定左边，依次从左遍历，两个指针移动，两个指针+固定的左边位置，3个数算0
func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			temp := nums[i] + nums[start] + nums[end]
			if temp == 0 {
				result = append(result, []int{nums[i], nums[start], nums[end]})
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				start++
				for start < end && nums[end] == nums[end-1] {
					end--
				}
				end--
			} else if temp > 0 {
				for start < end && nums[end] == nums[end-1] {
					end--
				}
				end--
			} else {
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				start++
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}

func Test_threeSum(t *testing.T) {
	//a := []int{-1, 0, 1, 2, -1, -4}
	//fmt.Println(threeSum(a))
}
