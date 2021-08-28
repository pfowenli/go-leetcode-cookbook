package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	node := head.Next.Next
	oddTail := head
	evenHead := head.Next
	evenTail := head.Next

	oddTail.Next = nil
	evenHead.Next = nil

	oddOrNot := false

	for node != nil {
		temp := node
		node = node.Next
		temp.Next = nil

		oddOrNot = !oddOrNot

		if oddOrNot {
			oddTail.Next = temp
			oddTail = temp
			continue
		}

		evenTail.Next = temp
		evenTail = temp
	}

	oddTail.Next = evenHead

	return head
}
