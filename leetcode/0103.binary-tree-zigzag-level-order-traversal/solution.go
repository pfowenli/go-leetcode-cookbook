package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	results := [][]int{}
	reverse := false

	for len(queue) != 0 {
		size := len(queue)
		list := []int{}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			list = append(list, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if reverse {
			for i := 0; i < len(list)/2; i++ {
				list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]
			}
		}
		reverse = !reverse
		results = append(results, list)
	}

	return results
}
