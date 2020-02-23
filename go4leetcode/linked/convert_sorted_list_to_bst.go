package linked

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	slowPoint := head
	fastPoint := head
	lastPoint := slowPoint
	for fastPoint.Next != nil && fastPoint.Next.Next != nil {
		lastPoint = slowPoint
		fastPoint = fastPoint.Next.Next
		slowPoint = slowPoint.Next
	}
	fastPoint = slowPoint.Next
	lastPoint.Next = nil
	currenNode := &TreeNode{Val: slowPoint.Val}
	if head != slowPoint {
		currenNode.Left = sortedListToBST(head)
	}
	currenNode.Right = sortedListToBST(fastPoint)
	return currenNode
}
