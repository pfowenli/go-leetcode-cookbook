package leetcode

func topKFrequent(nums []int, k int) []int {
	numFrequencyMap := make(map[int]int)
	for _, num := range nums {
		numFrequencyMap[num] = numFrequencyMap[num] + 1
	}

	buckets := make([][]int, len(nums)+1)
	for num, frequency := range numFrequencyMap {
		if buckets[frequency] == nil {
			buckets[frequency] = []int{}
		}
		buckets[frequency] = append(buckets[frequency], num)
	}

	topK := []int{}
	for i := len(buckets) - 1; i > 0 && len(topK) < k; i-- {
		if buckets[i] == nil {
			continue
		}
		if len(buckets[i]) < k-len(topK) {
			topK = append(topK, buckets[i]...)
			continue
		}
		topK = append(topK, buckets[i][0:k-len(topK)]...)
	}

	return topK
}
