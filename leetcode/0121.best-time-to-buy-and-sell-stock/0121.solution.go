package leetcode

func maxProfit(prices []int) int {
	minPrice := prices[0]
	result := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > minPrice {
			result = max(result, prices[i]-minPrice)
			continue
		}

		minPrice = prices[i]
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
