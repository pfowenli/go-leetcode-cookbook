package leetcode

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1] < intervals[b][1]
	})

	end := intervals[0][1]
	count := 1

	for _, interval := range intervals[1:] {
		if interval[0] < end {
			continue
		}

		end = interval[1]
		count++
	}

	return len(intervals) - count
}
