package leetcode

func trailingZeroes(n int) int {
	num := int(0)

	for n >= 5 {
		n /= 5
		num += n
	}

	return num
}
