package leetcode

import (
	"testing"
)

func TestFindDuplicateI(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	expected := int(2)

	if got := findDuplicate(nums); got != expected {
		t.Errorf("findDuplicate() = %d, expected = %d", got, expected)
	}
}

func TestFindDuplicateII(t *testing.T) {
	nums := []int{3, 1, 3, 4, 2}
	expected := int(3)

	if got := findDuplicate(nums); got != expected {
		t.Errorf("findDuplicate() = %d, expected = %d", got, expected)
	}
}

func TestFindDuplicateIII(t *testing.T) {
	nums := []int{1, 1}
	expected := int(1)

	if got := findDuplicate(nums); got != expected {
		t.Errorf("findDuplicate() = %d, expected = %d", got, expected)
	}
}

func TestFindDuplicateIV(t *testing.T) {
	nums := []int{1, 1, 2}
	expected := int(1)

	if got := findDuplicate(nums); got != expected {
		t.Errorf("findDuplicate() = %d, expected = %d", got, expected)
	}
}
