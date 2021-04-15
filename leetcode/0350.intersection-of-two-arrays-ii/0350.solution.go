package leetcode

func intersect(nums1 []int, nums2 []int) []int {
    countMap := make(map[int]int)
    for _, num := range nums1 {
        countMap[num] += 1        
    }
    
    var results []int
    for _, num :=range nums2 {
        if countMap[num] > 0 {
            results = append(results, num)
            countMap[num] -= 1
        }
    }
    
    return results
}

