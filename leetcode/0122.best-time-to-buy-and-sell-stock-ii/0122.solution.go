package leetcode

func maxProfit(prices []int) int {
	profit := int(0)

	for index, price := range prices {
		if index == 0 {
			continue
		}

		if price > prices[index-1] {
			profit += price - prices[index-1]
		}
	}

	return profit
}
