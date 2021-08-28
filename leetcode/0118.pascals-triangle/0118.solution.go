package leetcode

func generate(numRows int) [][]int {
	var results [][]int

	results = append(results, []int{1})

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		results = append(results, row)

		for j := 0; j < i; j++ {
			results[i][j] += results[i-1][j]
			results[i][j+1] += results[i-1][j]
		}
	}

	return results
}
