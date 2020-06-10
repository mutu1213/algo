package array

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {
	var result = math.MaxInt32
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	var minTemp = math.MaxInt32
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			temp := nums[i] + nums[start] + nums[end] - target
			absTemp := int(math.Abs(float64(temp)))
			if absTemp < minTemp {
				minTemp = absTemp
				result = nums[i] + nums[start] + nums[end]
			}
			if temp == 0 {
				return target
			} else if temp > 0 {
				end--
			} else {
				start++
			}

		}
	}
	return result
}

func Test_threeSumClosest(t *testing.T) {
	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 1, 2}, 3))
	//fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
