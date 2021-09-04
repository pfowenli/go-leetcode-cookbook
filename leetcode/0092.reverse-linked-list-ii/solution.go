package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	fakeHead := &ListNode{-1, head}
	leftTail := fakeHead

	for i := 1; i < left; i++ {
		leftTail = leftTail.Next
	}

	midHead := leftTail.Next
	midTail := leftTail.Next
	leftTail.Next = nil

	for i := left; i < right; i++ {
		midTail = midTail.Next
	}

	rightHead := midTail.Next
	midTail.Next = nil

	fakeMidHead := &ListNode{-1, midHead}
	fakeRightHead := &ListNode{-1, rightHead}

	for fakeMidHead.Next != nil {
		current := fakeMidHead.Next

		fakeMidHead.Next = current.Next

		current.Next = fakeRightHead.Next
		fakeRightHead.Next = current
	}

	leftTail.Next = fakeRightHead.Next

	return fakeHead.Next
}
