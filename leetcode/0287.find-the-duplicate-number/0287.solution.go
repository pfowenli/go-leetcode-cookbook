package leetcode

func findDuplicate(nums []int) int {
	result := int(0)
	numSet := make(map[int]bool)

	for _, num := range nums {
		if numSet[num] {
			result = num
			break
		}
		numSet[num] = true
	}

	return result
}
