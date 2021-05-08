package leetcode

func findKthLargest(nums []int, k int) int {
	nums = sortNums(nums)

	return nums[k-1]
}

func sortNums(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	results := []int{}

	pivot := len(nums) / 2
	leftNums := sortNums(nums[0:pivot])
	rightNums := sortNums(nums[pivot:])

	leftIndex := int(0)
	rightIndex := int(0)

	for {
		if leftIndex == len(leftNums) {
			results = append(results, rightNums[rightIndex:]...)
			break
		}
		if rightIndex == len(rightNums) {
			results = append(results, leftNums[leftIndex:]...)
			break
		}

		if leftNums[leftIndex] > rightNums[rightIndex] {
			results = append(results, leftNums[leftIndex])
			leftIndex++
			continue
		}

		results = append(results, rightNums[rightIndex])
		rightIndex++
	}

	return results
}
