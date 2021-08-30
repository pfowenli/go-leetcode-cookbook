package leetcode

func rob(nums []int) int {
	if len(nums) <= 3 {
		return maxElement(nums)
	}

	buckets := make([]int, len(nums)-1)

	for idx, num := range nums[:len(nums)-1] {
		pre := 0
		if idx-2 >= 0 {
			pre = max(pre, buckets[idx-2])
		}
		if idx-3 >= 0 {
			pre = max(pre, buckets[idx-3])
		}

		buckets[idx] = num + pre
	}

	result := max(buckets[len(buckets)-2], buckets[len(buckets)-1])

	for idx, num := range nums[1:len(nums)] {
		pre := 0
		if idx-2 >= 0 {
			pre = max(pre, buckets[idx-2])
		}
		if idx-3 >= 0 {
			pre = max(pre, buckets[idx-3])
		}

		buckets[idx] = num + pre
	}

	result = max(result, buckets[len(buckets)-2])
	result = max(result, buckets[len(buckets)-1])

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxElement(nums []int) int {
	result := nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] > result {
			result = nums[idx]
		}
	}
	return result
}
