package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int)

	for idx, num2 := range nums2 {
		numMap[num2] = idx
	}

	results := []int{}

	for _, num1 := range nums1 {
		result := -1

		for j := numMap[num1] + 1; j < len(nums2); j++ {
			if num1 < nums2[j] {
				result = nums2[j]
				break
			}
		}

		results = append(results, result)
	}

	return results
}
