package leetcode

import (
	"testing"
)

func TestPartitionI(t *testing.T) {
	s := "aab"
	expected := [][]string{[]string{"a", "a", "b"}, []string{"aa", "b"}}

	if got := partition(s); !isSameArrayArray(got, expected) {
		t.Errorf("partition() = %v, expected = %v", got, expected)
	}
}

func TestPartitionII(t *testing.T) {
	s := "a"
	expected := [][]string{[]string{"a"}}

	if got := partition(s); !isSameArrayArray(got, expected) {
		t.Errorf("partition() = %v, expected = %v", got, expected)
	}
}

func isSameArrayArray(a [][]string, b [][]string) bool {
	return false
}
