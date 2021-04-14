package leetcode

func firstUniqChar(s string) int {
    countMap := make(map[int32]int)
    indexMap := make(map[int32]int)
    
    for index, letter := range s {
        countMap[letter] += 1

        if countMap[letter] == 1 {
            indexMap[letter] = index
            continue
        }
        
        delete (indexMap, letter)
    }
    
    result := -1

    for _, value := range indexMap {
        if result > value || result == -1 {
            result = value
        }
    }
    
    return result
}

