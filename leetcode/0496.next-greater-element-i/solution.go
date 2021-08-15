package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int)
	stack := []int{}

	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			numMap[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, num)
		numMap[num] = -1
	}

	for idx, num := range nums1 {
		nums1[idx] = numMap[num]
	}

	return nums1
}
