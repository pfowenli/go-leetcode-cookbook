package leetcode

import (
	"testing"
)

func TestStrStrI(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	expected := 2
	if got := strStr(haystack, needle); got != expected {
		t.Errorf("strStr() = %d, expected %d", got, expected)
	}
}

func TestStrStrII(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"
	expected := -1
	if got := strStr(haystack, needle); got != expected {
		t.Errorf("strStr() = %d, expected %d", got, expected)
	}
}

func TestStrStrIII(t *testing.T) {
	haystack := ""
	needle := ""
	expected := 0
	if got := strStr(haystack, needle); got != expected {
		t.Errorf("strStr() = %d, expected %d", got, expected)
	}
}
