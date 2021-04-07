package leetcode

func majorityElement(nums []int) int {
    var maxCount int
    var targetNum int
    var countHash = make(map[int]int)
  
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        count := countHash[num] + 1
        countHash[num] = count

        if count > maxCount {
            targetNum = num
            maxCount = count
        }
    }

    return targetNum
}
