package leetcode

func subsets(nums []int) [][]int {
	results := [][]int{[]int{}}

	for _, num := range nums {
		combinations := [][]int{}

		for _, combination := range results {
			combinations = append(combinations, combination)
			combinations = append(combinations, temp := append(combination, int(num)))
		}

		results = combinations
	}

	return results
}
