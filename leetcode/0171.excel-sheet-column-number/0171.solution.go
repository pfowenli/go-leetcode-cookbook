package leetcode

func titleToNumber(columnTitle string) int {
    num := int(0)
    
    for _, c := range columnTitle {
        num = num * 26 + int(c) - 64
    }
    
    return num
}

