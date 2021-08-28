package leetcode

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for index := 0; index <= len(haystack)-len(needle); index++ {
		if isSameString(haystack[index:index+len(needle)], needle) {
			return index
		}
	}
	return -1
}

func isSameString(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for index, char := range a {
		if string(char) != b[index:index+1] {
			return false
		}
	}
	return true
}
