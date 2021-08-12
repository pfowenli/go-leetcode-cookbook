package leetcode

func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func (a, b int) bool {
        return points[a][0] < points[b][0]
    })
    
    end := points[0][1]
    count := 1
    
    for i := 1; i < len(points); i++ {
        if points[i][0] > end {
            end = points[i][1]
            count++
            continue
        }
        
        end = min(end, points[i][1])
    }
    
    return count
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

