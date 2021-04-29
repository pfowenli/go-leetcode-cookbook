package leetcode

import (
    "strings"
)

func isPalindrome(s string) bool {    
    left := int(-1)
    right := len(s)
    
    for {
        leftChar := ""
        rightChar := ""

        for {
            left++
            if left >= len(s) {
                break
            }

            leftChar = strings.ToLower(string(s[left]))
            if isAlphaNumeric(leftChar) {
                break
            }
        }

        for {
            right--
            if right <= 0 {
                break
            }

            rightChar = strings.ToLower(string(s[right]))
            if isAlphaNumeric(rightChar) {
                break
            }
        }
        
        if left >= right {
            break
        }
        
        if leftChar != rightChar {
            return false
        }
    }
    
    return true    
}

func isAlphaNumeric(char string) bool {
    if char >= "a" && char <= "z" {
        return true
    }
    if char >= "0" && char <= "9" {
        return true
    }    
    return false
}

:
