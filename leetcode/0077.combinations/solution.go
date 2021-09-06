package leetcode

func combine(n int, k int) [][]int {
	results := [][]int{}
	nums := []int{}

	var backtrack func(base int)
	backtrack = func(base int) {

		if len(nums) == k {
			result := append([]int{}, nums...)
			results = append(results, result)
			return
		}

		for val := base + 1; val <= n; val++ {
			nums = append(nums, val)
			backtrack(val)
			nums = nums[:len(nums)-1]
		}
	}

	backtrack(0)

	return results
}
