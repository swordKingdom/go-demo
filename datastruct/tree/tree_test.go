package tree

import "testing"

func TestCreatBinaryTree(t *testing.T) {
	arr := []int{1, 3, 4}
	root := CreatBinaryTree(arr)
	BinaryTreeTreePrint(root)
}
