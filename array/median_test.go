package array

//查找两个有序数组中的中位数,O(m+n)解法
//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	key := (len(nums1) + len(nums2) + 1) / 2
	key1 := (len(nums1) + len(nums2) + 2) / 2
	var result []int
	if len(nums1) == 0 {
		return float64(nums2[key-1]+nums2[key1-1]) / 2
	}
	if len(nums2) == 0 {
		return float64(nums1[key-1]+nums1[key1-1]) / 2
	}
	for i+j < key1 {
		if i < len(nums1) {
			if j < len(nums2) {
				if nums1[i] < nums2[j] {
					result = append(result, nums1[i])
					i++
				} else {
					result = append(result, nums2[j])
					j++
				}
			} else {
				result = append(result, nums1[i])
				i++
			}
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	return float64(result[key-1]+result[key1-1]) / 2
}
