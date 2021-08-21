package leetcode

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	results := [][]int{}

	for i, num := range nums {
		subNums := make([]int, len(nums)-1)
		for idx := 0; idx < i; idx++ {
			subNums[idx] = nums[idx]
		}
		for idx := i + 1; idx < len(nums); idx++ {
			subNums[idx-1] = nums[idx]
		}

		subPermutations := permute(subNums)

		results = append(results, make([][]int, len(subPermutations))...)

		for j, subPermutation := range subPermutations {
			results[i*len(subPermutations)+j] = append(subPermutation, num)
		}
	}

	return results
}
