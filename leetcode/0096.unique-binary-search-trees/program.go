package leetcode

func numTrees(n int) int {
	countList := make([]int, n+1)
	countList[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			countList[i] += countList[j] * countList[i-1-j]
		}
	}

	return countList[n]
}
