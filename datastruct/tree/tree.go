package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree interface {
	IsEmpty() bool
	Clear()
	Size() int
}
