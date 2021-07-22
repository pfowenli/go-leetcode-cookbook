package leetcode

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	directions := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 || board[i][j] != 'O' {
			return
		}

		board[i][j] = 'V'
		for _, d := range directions {
			dfs(i+d[0], j+d[1])
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'V' {
				board[i][j] = 'O'
				continue
			}
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
