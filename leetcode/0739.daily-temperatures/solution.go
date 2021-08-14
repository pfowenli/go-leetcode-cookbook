package leetcode

func dailyTemperatures(temperatures []int) []int {
    stack := []int{}
    results := make([]int, len(temperatures))
    
    for i, temperature := range temperatures {
        j := len(stack) - 1

        for j >= 0 {
            if temperatures[stack[j]] >= temperature {
                break
            }

            results[stack[j]] = i - stack[j]
            j--
        }

        stack = append(stack[:j + 1], i)
    }
    
    return results  
}

