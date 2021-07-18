package leetcode

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(a, b int) bool {
        if intervals[a][0] > intervals[b][0] {
            return false
        }
        if intervals[a][0] == intervals[b][0] && intervals[a][1] < intervals[b][1] {
            return false
        }
        return true
    })
    
    res := [][]int{}
    
    start := intervals[0][0]
    end := intervals[0][1]
    
    for _, interval := range intervals {
        if end < interval[0] {
            res = append(res, []int{start, end})
            start = interval[0]
            end = interval[1]
            continue
        }
        
        if end < interval[1] {
            end = interval[1]
        }
    }
    
    res = append(res, []int{start, end})
    
    return res
}
