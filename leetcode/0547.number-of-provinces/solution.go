package leetcode

func findCircleNum(isConnected [][]int) int {
	result := 0

	var helper func(i int)
	helper = func(i int) {
		if isConnected[i][i] == 0 {
			return
		}

		for j := range isConnected[0] {
			if isConnected[i][j] == 1 {
				isConnected[i][j] = 0
				helper(j)
			}
		}
	}

	for i := range isConnected {
		if isConnected[i][i] == 1 {
			helper(i)
			result++
		}
	}

	return result
}
