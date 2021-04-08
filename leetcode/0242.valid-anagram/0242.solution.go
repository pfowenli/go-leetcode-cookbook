package leetcode

func isAnagram(s string, t string) bool {
    var countMap = make(map[string]int)

    for i := 0; i < len(s); i++ {
        letter := s[i : i + 1]
        count := countMap[letter]
        countMap[letter] = count + 1
    }
    
    for i := 0; i < len(t); i++ {
        letter := t[i : i + 1]
        count := countMap[letter]

        if count == 0 {
            return false
        }

        countMap[letter] = count - 1
    }
    
    for _, count := range countMap {
        if count != 0 {
            return false
        }
    }

    return true
}

