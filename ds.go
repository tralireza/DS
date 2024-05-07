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

// 2816m Double a Number Represented as a Linked List
func doubleIt(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	l := doubleIt(head.Next)

	head.Val *= 2
	if l != head.Next {
		head.Val++
	}

	if head.Val >= 10 {
		n := &ListNode{head.Val / 10, head}
		head.Val %= 10
		head = n
	}
	return head
}
