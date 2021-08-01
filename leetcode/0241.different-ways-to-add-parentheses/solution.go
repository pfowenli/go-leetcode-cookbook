package leetcode

import "strconv"

func diffWaysToCompute(expression string) []int {
	results := []int{}

	for idx := range expression {
		if expression[idx] == '+' || expression[idx] == '-' || expression[idx] == '*' {
			aList := diffWaysToCompute(expression[:idx])
			bList := diffWaysToCompute(expression[idx+1:])

			for _, a := range aList {
				for _, b := range bList {
					value := 0
					switch expression[idx] {
					case '+':
						value = a + b
						break
					case '-':
						value = a - b
						break
					case '*':
						value = a * b
						break
					case '/':
						value = a / b
						break
					}

					results = append(results, value)
				}
			}
		}
	}

	if len(results) == 0 {
		value, _ := strconv.Atoi(expression)
		return []int{value}
	}

	return results
}
