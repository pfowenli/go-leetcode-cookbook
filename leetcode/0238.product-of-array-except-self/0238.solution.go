package leetcode

func productExceptSelf(nums []int) []int {
    products := make([]int, len(nums))
    forwardProducts := make([]int, len(nums))
    backwardProducts := make([]int, len(nums))
    
    forwardProducts[0] = 1
    for i := 1; i < len(nums); i++ {
        forwardProducts[i] = forwardProducts[i - 1] * nums[i - 1]        
    }
    
    backwardProducts[len(nums) - 1] = 1
    for i := len(nums) - 2; i >= 0; i-- {
        backwardProducts[i] = backwardProducts[i + 1] * nums[i + 1]
    }
    
    for i := 0; i < len(nums); i++ {
        products[i] = forwardProducts[i] * backwardProducts[i]
    }

    return products
}

