package DS

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2487m Remove Nodes from Linked List
func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeNodes(head.Next)
	if head.Next != nil && head.Val < head.Next.Val {
		head = head.Next
	}
	return head
}
