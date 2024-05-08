package DS

import (
	"slices"
	"strconv"
)

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

// 506 Relative Ranks
func findRelativeRanks(score []int) []string {
	pQueue := make([][]int, len(score))
	for i, v := range score {
		pQueue[i] = []int{i, v}
	}
	slices.SortFunc(pQueue, func(x, y []int) int { return y[1] - x[1] })

	Rank := make([]string, len(score))
	for i, v := range pQueue {
		var rank string
		switch i {
		case 0:
			rank = "Gold Medal"
		case 1:
			rank = "Silver Medal"
		case 2:
			rank = "Bronze Medal"
		default:
			rank = strconv.Itoa(i + 1)
		}
		Rank[v[0]] = rank
	}
	return Rank
}
