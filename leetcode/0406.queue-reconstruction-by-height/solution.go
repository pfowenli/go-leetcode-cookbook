package leetcode

import "sort"

func reconstructQueue(people [][]int) [][]int {
    sort.Slice(people, func(a, b int) bool {
        if people[a][0] > people[b][0] {
            return true
        }
        if people[a][0] == people[b][0] && people[a][1] < people[b][1] {
            return true
        }
        return false
    })
    
    results := [][]int{}
    
    for _, p := range people {
        results = append(results, []int{})
        copy(results[p[1]+1:], results[p[1]:])
        results[p[1]] = p
    }
    
    return results
}

