package leetcode

func climbStairs(n int) int {
    countMap := make(map[int]int)
    countMap[1] = 1
    countMap[2] = 2
    
    for i := 3; i <= n; i++ {
        countMap[i] = countMap[i - 2] + countMap[i - 1]
    }

    return countMap[n]
}
