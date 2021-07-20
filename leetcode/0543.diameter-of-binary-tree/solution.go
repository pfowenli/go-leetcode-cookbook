package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0

	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := getDepth(node.Left)
		r := getDepth(node.Right)

		if l+r > result {
			result = l + r
		}

		if l > r {
			return l + 1
		}
		return r + 1
	}

	getDepth(root)

	return result
}
