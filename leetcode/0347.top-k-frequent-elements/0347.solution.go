package leetcode

func topKFrequent(nums []int, k int) []int {
	maxCount := int(0)
	countMap := make(map[int]int)
	numSetMap := make(map[int](map[int]bool))

	for _, num := range nums {
		count := countMap[num]

		if count == maxCount {
			maxCount += 1
		}

		countMap[num] += 1

		if numSetMap[count + 1] == nil {
			numSetMap[count + 1] = make(map[int]bool)
		}
		numSetMap[count + 1][num] = true

		if count > 0 {
			delete(numSetMap[count], num)
		}
	}

	results := []int{}

	count := maxCount
	for count >= 0 {
		numSet := numSetMap[count]

		for num, _ := range numSet {
			results = append(results, num)
			if len(results) == k {
				return results
			}
		}

		count -= 1
	}

	return results
}
