package leetcode

func subsets(nums []int) [][]int {
	results := [][]int{}
	results = append(results, []int{})

	for _, num := range nums {
		for _, result := range results {
			combination := []int{}
			combination = append(combination, result...)
			combination = append(combination, num)

			results = append(results, combination)
		}
	}

	return results
}
