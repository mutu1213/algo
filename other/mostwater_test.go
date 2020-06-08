package other
//leetcode https://leetcode-cn.com/problems/container-with-most-water/

func maxArea(height []int) int {
	var result int
	i,j:=0,len(height)-1
	for i<j {
		temp := 0
		//当两个相等的时候，不需要额外考虑，因为如果里面有一个比外面一层小，那么肯定是外面这层大，里面两个都比外面一层大的时候，移动任意一个都行
		if height[i]<height[j]{
			i++
			temp = height[i] * (j-i)
		} else  {
			j--
			temp = height[j] * (j-i)
		}
		if temp> result {
			result = temp
		}
	}
	return result
}
