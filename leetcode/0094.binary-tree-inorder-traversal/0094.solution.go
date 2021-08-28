package leetcode

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := []int{}

	nums = append(nums, inorderTraversal(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inorderTraversal(root.Right)...)

	return nums
}
