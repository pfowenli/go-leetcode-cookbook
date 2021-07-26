package leetcode

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	m := len(grid1)
	n := len(grid1[0])

	count := 0
	isSubIsland := false

	var dfs func(r, s int)
	dfs = func(r, s int) {
		if r < 0 || r > m-1 || s < 0 || s > n-1 || grid2[r][s] == 0 {
			return
		}

		grid2[r][s] = 0

		if grid1[r][s] == 0 {
			isSubIsland = false
		}

		for _, d := range directions {
			dfs(r+d[0], s+d[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				isSubIsland = true

				dfs(i, j)

				if isSubIsland {
					count++
				}
			}
		}
	}

	return count
}
