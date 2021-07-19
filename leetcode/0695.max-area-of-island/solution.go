package leetcode

func maxAreaOfIsland(grid [][]int) int {
	directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	m, n := len(grid), len(grid[0])
	result := 0

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i > m-1 || j < 0 || j > n-1 || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0
		sum := 1

		for _, d := range directions {
			sum += dfs(i+d[0], j+d[1])
		}

		return sum
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}

			sum := dfs(i, j)

			if result < sum {
				result = sum
			}
		}
	}

	return result

}
