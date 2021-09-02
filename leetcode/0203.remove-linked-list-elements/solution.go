package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	fakeHead := &ListNode{-1, head}
	node := fakeHead

	for node.Next != nil {
		if node.Next.Val != val {
			node = node.Next
		} else {
			node.Next = node.Next.Next
		}
	}

	return fakeHead.Next
}
