package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"math"
)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftNode, rightNode := root.Left, root.Right
	leftHeight, rightHeight := 0, 0

	for leftNode != nil {
		leftHeight++
		leftNode = leftNode.Left
	}

	for rightNode != nil {
		rightHeight++
		rightNode = rightNode.Right
	}

	if leftHeight == rightHeight {
		return int(math.Pow(2, float64(leftHeight+1)) - 1)
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
