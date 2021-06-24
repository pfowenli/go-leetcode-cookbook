package leetcode

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	var rev *ListNode = nil

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		previous := rev
		rev = slow
		slow = slow.Next
		rev.Next = previous
	}

	if fast != nil {
		slow = slow.Next
	}

	for rev != nil && slow != nil {
		if slow.Val != rev.Val {
			return false
		}
		slow = slow.Next
		rev = rev.Next
	}

	return true
}
