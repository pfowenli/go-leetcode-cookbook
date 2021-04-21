package leetcode

func plusOne(digits []int) []int {
    for index := len(digits) - 1; index >= 0; index-- {
        digits[index] += 1

        if digits[index] != 10 {
            break
        }

        digits[index] = 0
    }

    if digits[0] == 0 {
        digits = make([]int, len(digits) + 1)
        digits[0] = 1
    }
    return digits      
}

