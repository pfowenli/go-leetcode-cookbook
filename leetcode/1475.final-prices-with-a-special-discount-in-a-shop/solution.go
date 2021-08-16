package main

import (
	"fmt"
)

func finalPrices(prices []int) []int {
    stack := []int{}
	results := make([]int, len(prices))

	for idx, price := range prices {
		for len(stack) > 0 && prices[stack[len(stack) - 1]] >= price {
			results[stack[len(stack) - 1]] = prices[stack[len(stack) - 1]] - price
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, idx)
		results[idx] = price
	}

	return results
}
