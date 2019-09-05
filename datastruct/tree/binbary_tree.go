package tree

import "fmt"

type BinaryTree struct {
	root *TreeNode
}

func binaryTreeTreePrint(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	binaryTreeTreePrint(root.Left)
	binaryTreeTreePrint(root.Right)
}

func BinaryTreeTreePrint(root *BinaryTree) {
	if root == nil {
		return
	}
	binaryTreeTreePrint(root.root)
}

func buildTree(root *TreeNode, arr []int, index int, arrLen int) {
	if index > arrLen {
		return
	}
	left := index*2 + 1
	right := left + 1
	if right > arrLen || left > arrLen {
		return
	}
	if left < arrLen {
		root.Left = &TreeNode{Val: arr[left]}
		buildTree(root.Left, arr, left, arrLen)
	}
	if right < arrLen {
		root.Right = &TreeNode{Val: arr[right]}
		buildTree(root.Right, arr, right, arrLen)
	}
}

func CreatBinaryTree(arr []int) *BinaryTree {
	arrLen := len(arr)
	if arrLen == 0 {
		return nil
	}
	res := &TreeNode{Val: arr[0]}
	buildTree(res, arr, 0, len(arr))
	return &BinaryTree{root: res}
}
