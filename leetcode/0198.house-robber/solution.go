package leetcode

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) >= 3 {
		nums[2] += nums[0]
	}

	for idx := 3; idx < len(nums); idx++ {
		nums[idx] += max(nums[idx-3], nums[idx-2])
	}

	return max(nums[len(nums)-2], nums[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
