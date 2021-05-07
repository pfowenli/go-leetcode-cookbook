package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNode(node *ListNode) {

	if node == nil {
		return
	}

	for {
		if node.Next == nil {
			break
		}
		node.Val = node.Next.Val

		if node.Next.Next == nil {
			break
		}

		node = node.Next
	}

	node.Next = nil
}
