package leetcode

func missingNumber(nums []int) int {
	length := len(nums)
	sum := (1 + length) * length / 2

	for _, value := range nums {
		sum -= value
	}

	return sum
}
