package leetcode

import (
	"sort"
	"testing"
)

func TestTopKFrequentI(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := int(2)
	expected := []int{1, 2}

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
