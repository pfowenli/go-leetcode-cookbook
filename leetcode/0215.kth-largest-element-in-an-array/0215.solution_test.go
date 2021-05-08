package leetcode

import (
	"testing"
)

func TestFindKthLargestI(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := int(2)
	expected := int(5)

	if got := findKthLargest(nums, k); got != expected {
		t.Errorf("findKthLargest() = %d, expected = %d", got, expected)
	}
}

func TestFindKthLargestII(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := int(4)
	expected := int(4)

	if got := findKthLargest(nums, k); got != expected {
		t.Errorf("findKthLargest() = %d, expected = %d", got, expected)
	}
}
