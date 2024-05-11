package DS

import (
	"log"
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

// 786 K-th Smallest Prime Fraction
func kthSmallestPrimeFraction(arr []int, k int) []int {
	Q := [][]int{}

	var Heapify func(int)
	Heapify = func(i int) {
		l, r, p := 2*i+1, 2*i+2, i
		if l < len(Q) && Q[l][0]*Q[p][1] < Q[l][1]*Q[p][0] {
			p = l
		}
		if r < len(Q) && Q[r][0]*Q[p][1] < Q[r][1]*Q[p][0] {
			p = r
		}
		if p != i {
			Q[i], Q[p] = Q[p], Q[i]
			Heapify(p)
		}
	}

	Init := func() {
		for i := len(Q); i >= 0; i-- {
			Heapify(i)
		}
	}

	Pop := func() []int {
		v := Q[0]
		Q[0] = Q[len(Q)-1]
		Q = Q[:len(Q)-1]
		Heapify(0)
		return v
	}

	Push := func(v []int) {
		Q = append(Q, v)
		i := len(Q) - 1
		for i > 0 && Q[i][0]*Q[(i-1)/2][1] < Q[i][1]*Q[(i-1)/2][0] {
			Q[i], Q[(i-1)/2] = Q[(i-1)/2], Q[i]
			i = (i - 1) / 2
		}
	}

	for i := range arr {
		Q = append(Q, []int{arr[i], arr[len(arr)-1], i, len(arr) - 1})
	}
	Init()

	log.Print("::", Q)

	for range k - 1 {
		log.Print(Q)

		v := Pop()
		n, d := v[2], v[3]-1
		if d > n {
			Push([]int{arr[n], arr[d], n, d})
		}
	}

	v := Q[0]
	return []int{v[0], v[1]}
}
