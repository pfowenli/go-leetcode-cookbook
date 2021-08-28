package leetcode

import (
	"testing"
)

func TestOddEventListI(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := []int{1, 3, 5, 2, 4}

	head := convertArrayToList(input)
	expected := convertArrayToList(output)

	if got := oddEvenList(head); !isSameList(got, expected) {
		t.Errorf("oddEvenList() = %v, expected = %v", convertListToArray(got), convertListToArray(expected))
	}
}

func TestOddEventListII(t *testing.T) {
	input := []int{2, 1, 3, 5, 6, 4, 7}
	output := []int{2, 3, 6, 7, 1, 5, 4}

	head := convertArrayToList(input)
	expected := convertArrayToList(output)

	if got := oddEvenList(head); !isSameList(got, expected) {
		t.Errorf("oddEvenList() = %v, expected = %v", convertListToArray(got), convertListToArray(expected))
	}
}

func convertArrayToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return (*ListNode)(nil)
	}

	head := &ListNode{nums[0], nil}

	node := head
	for _, num := range nums[1:] {
		node.Next = &ListNode{num, nil}
		node = node.Next
	}

	return head
}

func convertListToArray(head *ListNode) []int {
	nums := []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return nums
}

func isSameList(a *ListNode, b *ListNode) bool {
	for {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}

		a = a.Next
		b = b.Next
	}

	return true
}
