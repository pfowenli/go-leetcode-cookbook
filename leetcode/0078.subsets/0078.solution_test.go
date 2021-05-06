package leetcode

import (
	"testing"
        "reflect"
)

func TestSubsetsI(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{[]int{}, []int{1}, []int{1, 2}, []int{1, 2, 3}, []int{2}, []int{2, 3}, []int{3}}

	if got := subsets(nums); !reflect.DeepEqual(got, expected) {
		t.Errorf("subsets() = %d, expected = %d", got, expected)
	}
}

func TestSubsetsII(t *testing.T) {
	nums := []int{1}
	expected := [][]int{[]int{}, []int{1}}

	if got := subsets(nums); !reflect.DeepEqual(got, expected) {
		t.Errorf("subsets() = %d, expected = %d", got, expected)
	}
}

func TestSubsetsIII(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	expected := [][]int{[]int{}, []int{1}, []int{1, 2}, []int{1, 2, 3}, []int{2}, []int{2, 3}, []int{3}}

	if got := subsets(nums); !reflect.DeepEqual(got, expected) {
		t.Errorf("subsets() = %d, expected = %d", got, expected)
	}
}
