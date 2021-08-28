package leetcode

func hammingWeight(num uint32) int {
	onesCount := 0

	for num != 0 {
		if num&1 == 1 {
			onesCount++
		}

		num = num >> 1
	}

	return onesCount
}
