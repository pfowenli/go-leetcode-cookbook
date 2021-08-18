package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	tempHead := &TreeNode{-1, nil, nil}
	tail := tempHead

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if last == nil {
			continue
		}

		if last.Left == nil && last.Right == nil {
			tail.Right = last
			tail = tail.Right
			continue
		}

		stack = append(stack, last.Right)
		last.Right = nil

		stack = append(stack, last)

		stack = append(stack, last.Left)
		last.Left = nil
	}

	return tempHead.Right
}
