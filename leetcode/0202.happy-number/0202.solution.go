package leetcode

func isHappy(n int) bool {
	valueMap := make(map[int]bool)

	for !valueMap[n] {
		valueMap[n] = true
		n = convertNumber(n)

		if n == 1 {
			return true
		}
	}

	return false
}

func convertNumber(num int) int {
	sum := 0

	for num > 0 {
		remain := num % 10
		sum += remain * remain
		num = num / 10
	}

	return sum
}
