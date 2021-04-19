package leetcode

func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }

    counter := 1
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] == nums[i] {
            continue
        }
        nums[counter] = nums[i]
        counter++
    }
    return counter
}

