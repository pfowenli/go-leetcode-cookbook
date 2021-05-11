package leetcode

import (
	"testing"
)

func TestUniquePathsI(t *testing.T) {
	m := int(3)
	n := int(7)
	expected := int(28)

	if got := uniquePaths(m, n); got != expected {
		t.Errorf("uniquePaths() = %d, expected = %d", got, expected)
	}
}

func TestUniquePathsII(t *testing.T) {
	m := int(3)
	n := int(2)
	expected := int(3)

	if got := uniquePaths(m, n); got != expected {
		t.Errorf("uniquePaths() = %d, expected = %d", got, expected)
	}
}

func TestUniquePathsIII(t *testing.T) {
	m := int(51)
	n := int(9)
	expected := int(1916797311)

	if got := uniquePaths(m, n); got != expected {
		t.Errorf("uniquePaths() = %d, expected = %d", got, expected)
	}
}
