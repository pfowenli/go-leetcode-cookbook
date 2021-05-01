package leetcode

import (
	"testing"
)

func TestInorderTraversalI(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	expected := []int{1, 3, 2}

	if got := inorderTraversal(root); !isSameIntArray(got, expected) {
		t.Errorf("inorderTraversal() = %v, expected %v", got, expected)
	}
}

func TestInorderTraversalII(t *testing.T) {
	var root *TreeNode
	expected := []int{}

	if got := inorderTraversal(root); !isSameIntArray(got, expected) {
		t.Errorf("inorderTraversal() = %v, expected %v", got, expected)
	}
}

func TestInorderTraversalIII(t *testing.T) {
	root := &TreeNode{1, nil, nil}
	expected := []int{1}

	if got := inorderTraversal(root); !isSameIntArray(got, expected) {
		t.Errorf("inorderTraversal() = %v, expected %v", got, expected)
	}
}

func TestInorderTraversalIV(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	expected := []int{2, 1}

	if got := inorderTraversal(root); !isSameIntArray(got, expected) {
		t.Errorf("inorderTraversal() = %v, expected %v", got, expected)
	}
}

func TestInorderTraversalV(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	expected := []int{1, 2}

	if got := inorderTraversal(root); !isSameIntArray(got, expected) {
		t.Errorf("inorderTraversal() = %v, expected %v", got, expected)
	}
}

func isSameIntArray(intArray1 []int, intArray2 []int) bool {
	if len(intArray1) != len(intArray2) {
		return false
	}

	for index, int1 := range intArray1 {
		if int1 != intArray2[index] {
			return false
		}
	}
	return true
}
