package leetcode

func moveZeroes(nums []int)  {
    count := 0
    temp := make([]int, len(nums))
    
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            count++
            continue
        }

        temp[i - count] = nums[i]
    }
    
    for i := 0; i < len(nums); i++ {
        nums[i] = temp[i]
    }
}

