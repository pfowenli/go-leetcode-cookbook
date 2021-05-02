package

func groupAnagrams(strs []string) [][]string {
    anagramGroupsMap := make(map[int][][]string)
    

    for _, str := range strs {
        anagramGroups := anagramGroupsMap[len(str)]

        inserted := false
        for index, anagrams := range anagramGroups {
            if isAnagram(anagrams[0], str) {
                anagramGroups[index] = append(anagrams, str)
                inserted = true
                break
            }
        }
        if inserted == false {
            anagramGroups = append(anagramGroups, []string{str})
        }
        
        anagramGroupsMap[len(str)] = anagramGroups
    }
    
    results := [][]string{}
    
    for _, anagramGroups := range anagramGroupsMap {
        results = append(results, anagramGroups...)
    }
    
    return results
}

func isAnagram(a string, b string) bool {
    letterCountMap := make(map[rune]int)
    
    for _, letter := range a {
        letterCountMap[letter] += 1
    }
    
    for _, letter := range b {
        if letterCountMap[letter] == 0 {
            return false
        }
        letterCountMap[letter] -= 1
    }
    
    for _, count := range letterCountMap {
        if count != 0 {
            return false
        }
    }
    
    return true
}

