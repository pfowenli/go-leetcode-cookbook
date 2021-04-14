package leetcode

func firstUniqChar(s string) int {
    countMap := make(map[int32]int)
    
    for _, letter := range s {
        countMap[letter] += 1
    }
    
    for index, letter := range s {
        if countMap[letter] == 1 {
            return index
        }
    }
    
    return -1
}

