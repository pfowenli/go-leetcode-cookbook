package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
    length := len(nums)

    if length == 0 {
        return nil
    }

    mid := length / 2
    
    leftTree := sortedArrayToBST(nums[0:mid])
    rightTree := sortedArrayToBST(nums[mid+1:length])
    
    tree := TreeNode{nums[mid], leftTree, rightTree}
    return &tree
}

