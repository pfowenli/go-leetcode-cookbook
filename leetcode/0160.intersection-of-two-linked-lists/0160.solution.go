package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    listNodeMap := make(map[*ListNode]bool)
    
    for headA != nil {
        listNodeMap[headA] = true
        headA = headA.Next
    }
    
    for headB != nil {
        if listNodeMap[headB] {
            return headB
        }
        headB = headB.Next
    }

    return nil
}

