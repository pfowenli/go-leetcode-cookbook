package leetcode

func groupAnagrams(strs []string) [][]string {
	letterCountMapMap := make(map[string]map[rune]int)
	anagramGroupMap := make(map[string][]string)

	for _, str := range strs {
		letterCountMap := make(map[rune]int)
		for _, letter := range str {
			letterCountMap[letter] += 1
		}

		anagram := str
		for a, m := range letterCountMapMap {
			if len(a) != len(str) {
				continue
			}
			if !isAnagram(m, letterCountMap) {
				continue
			}
			anagram = a
			break
		}

		if anagram == str {
			letterCountMapMap[anagram] = letterCountMap
		}

		anagramGroupMap[anagram] = append(anagramGroupMap[anagram], str)
	}

	results := [][]string{}
	for _, anagramGroup := range anagramGroupMap {
		results = append(results, anagramGroup)
	}

	return results
}

func isAnagram(map1 map[rune]int, map2 map[rune]int) bool {
	for letter, count := range map1 {
		if map2[letter] != count {
			return false
		}
	}
	return true
}
