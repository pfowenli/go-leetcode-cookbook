package leetcode

import (
	"sort"
	"testing"
)

func TestTopKFrequentI(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 5, 5}
	k := int(4)
	expected := []int{1, 4, 2, 5}

	if got := topKFrequent(nums, k); !isSameSet(got, expected) {
		t.Errorf("topKFrequent() = %d, expected = %d", got, expected)
	}
}

func TestTopKFrequentII(t *testing.T) {
	nums := []int{1}
	k := int(1)
	expected := []int{1}

	if got := topKFrequent(nums, k); !isSameSet(got, expected) {
		t.Errorf("topKFrequent() = %d, expected = %d", got, expected)
	}
}

func isSameSet(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for index, num := range a {
		if num != b[index] {
			return false
		}
	}

	return true
}
