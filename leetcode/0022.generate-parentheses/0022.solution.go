package leetcode

import "strings"

func generateParenthesis(n int) []string {
    return combineParentheses([]string{}, []string{}, n)
}

func combineParentheses(parenthesesArray [] string, combinations []string, numberOfPairs int) []string {
    leftParenthesisCount := int(0)
    rightParenthesisCount := int(0)

    for _, c := range parenthesesArray {
        if c == "(" {
            leftParenthesisCount++
        }
    }
    
    rightParenthesisCount = len(parenthesesArray) - leftParenthesisCount

    if leftParenthesisCount > numberOfPairs {
        return combinations
    }
    
    if rightParenthesisCount > len(parenthesesArray) / 2 {
        return combinations
    }

    if len(parenthesesArray) == numberOfPairs * 2 {
        combinations = append(combinations, strings.Join(parenthesesArray, ""))
        return combinations
    }
    
    combinations = combineParentheses(append(parenthesesArray, "("), combinations, numberOfPairs)
    combinations = combineParentheses(append(parenthesesArray, ")"), combinations, numberOfPairs)
    
    return combinations
}

