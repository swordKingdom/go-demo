package tree

type AvlTreeNode struct {
	Parent *AvlTreeNode
	Val    interface{}
	Left   *AvlTreeNode
	Right  *AvlTreeNode
}

type AvlTree struct {
	root *AvlTreeNode
	size int
	m    int
}

func (a *AvlTree) IsEmpty() bool {
	return a.size == 0
}

func (a *AvlTree) Clear() {
	a.root = nil
	a.size = 0
	a.m = 0
}

func (a *AvlTree) Size() int {
	return a.size
}
