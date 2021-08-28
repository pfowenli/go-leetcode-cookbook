package leetcode

func singleNumber(nums []int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		// XOR operator
		result ^= nums[i]
	}

	return result
}
