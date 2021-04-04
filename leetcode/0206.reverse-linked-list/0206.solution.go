/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    previous := head
    current := head.Next
    head.Next = nil

    for current.Next != nil {
        next := current.Next
        current.Next = previous

        previous = current
        current = next
    }

    current.Next = previous

    return current
}

