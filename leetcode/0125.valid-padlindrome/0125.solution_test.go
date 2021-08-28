package leetcode

import (
	"testing"
)

func TestIsPalindromeI(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	expected := true
	if got := isPalindrome(s); got != expected {
		t.Errorf("isPalindrome() = %t, expected %t", got, expected)
	}
}

func TestIsPalindromeII(t *testing.T) {
	s := "race a car"
	expected := false
	if got := isPalindrome(s); got != expected {
		t.Errorf("isPalindrome() = %t, expected %t", got, expected)
	}
}
