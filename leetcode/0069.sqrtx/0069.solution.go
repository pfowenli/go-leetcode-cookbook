package leetcode

func mySqrt(x int) int {
    result := int(1)
    
    for result * result <= x {
        result++
    }
    
    result--
    
    return result
}

