package tree

type AvlTreeNode struct {
	Parent *AvlTreeNode
	Val    interface{}
	Left   *AvlTreeNode
	Right  *AvlTreeNode
	b      int8
}

type AvlTree struct {
	root       *AvlTreeNode
	size       int
	m          int
	comparator func(a, b interface{}) int8
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

func putFix(c int8, node **AvlTreeNode) bool {
	s := *node
	if s.b == 0 {
		s.b = c
		return true
	}

	if s.b == -c {
		s.b = 0
		return false
	}
	var b int8
	if (c+1)/2 == 0 {
		b = s.Left.b
	} else {
		b = s.Right.b
	}
	if b == c {
		//单旋

	} else {
		//双旋
	}
	*node = s
	return false
}

func (a *AvlTree) put(val interface{}, parent *AvlTreeNode, children **AvlTreeNode) bool {
	q := *children
	if q == nil {
		a.size++
		*children = &AvlTreeNode{Val: val, Parent: parent}
		return true
	}
	cmp := a.comparator(val, q.Val)
	if cmp == 0 {
		return false
	}
	var fix bool
	if cmp > 0 {
		fix = a.put(val, q, &q.Right)
	} else {
		fix = a.put(val, q, &q.Left)
	}
	if fix {
		return putFix(cmp, children)
	}
	return false
}

func (a *AvlTree) Put(val int) {

}
