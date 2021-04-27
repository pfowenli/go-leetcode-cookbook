package leetcode

func isPowerOfThree(n int) bool {
    if n < 1 {
        return false
    }
    
    product := int(1)
    for product <= n {
        if product == n {
            return true
        }
        product *= 3
    }
    
    return false
}

