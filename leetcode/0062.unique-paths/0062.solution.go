package leetcode

func uniquePaths(m int, n int) int {
	pathCounts := make([]int, n)
	for i := range pathCounts {
		pathCounts[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < len(pathCounts); j++ {
			pathCounts[j] += pathCounts[j-1]
		}
	}

	return pathCounts[n-1]
}
