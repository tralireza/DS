package DS

import (
	"log"
	"math"
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

// 2373 Largest Local Values in a Matrix
func largestLocal(grid [][]int) [][]int {
	Mx := [][]int{}

	for r := 1; r < len(grid)-1; r++ {
		Mx = append(Mx, make([]int, len(grid[r])-2))

		for c := 1; c < len(grid[r])-1; c++ {
			x := grid[r][c]
			for i := r - 1; i <= r+1; i++ {
				for j := c - 1; j <= c+1; j++ {
					x = max(x, grid[i][j])
				}
			}
			Mx[r-1][c-1] = x
		}
	}

	return Mx
}

// 1219m Path with Maximum Gold
func getMaximumGold(grid [][]int) int {
	Rows, Cols := len(grid), len(grid[0])
	gold := 0

	dirs := []int{0, 1, 0, -1, 0}

	var Search func(r, c, g int) int
	Search = func(r, c, g int) int {
		log.Printf("G: %2d  (%d,%d) -> %v", g, r, c, grid)
		for i := range dirs[:4] {
			x, y := r+dirs[i], c+dirs[i+1]
			if x >= 0 && x < Rows && y >= 0 && y < Cols && grid[x][y] > 0 {
				g += grid[x][y]
				grid[x][y] *= -1
				gold = max(Search(x, y, g), gold)

				grid[x][y] *= -1
				g -= grid[x][y]
			}
		}
		return g
	}

	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if grid[r][c] > 0 {
				grid[r][c] *= -1
				gold = max(Search(r, c, -grid[r][c]), gold)
				grid[r][c] *= -1
			}
		}
	}

	return gold
}

// 861m Score After Flipping Matrix
func matrixScore(grid [][]int) int {
	for r := 0; r < len(grid); r++ {
		if grid[r][0] != 1 {
			for c := 0; c < len(grid[r]); c++ {
				grid[r][c] = ^(grid[r][c] & 1) & 1
			}
		}
	}
	log.Print(" -> ", grid)

	for c := 1; c < len(grid[0]); c++ {
		ones := 0
		for r := 0; r < len(grid); r++ {
			ones += grid[r][c]
		}

		if ones < len(grid)-ones {
			for r := 0; r < len(grid); r++ {
				grid[r][c] = ^(grid[r][c] & 1) & 1
			}
		}
	}
	log.Print(" -> ", grid)

	score := 0
	for r := 0; r < len(grid); r++ {
		v := 0
		for c := 0; c < len(grid[0]); c++ {
			v += v + grid[r][c]
		}
		score += v
	}
	return score
}

// 857h Minimum Cost to Hire K Workers
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	Q := []int{} // Heap: Priority Queue

	var Heapify func(i int)
	Heapify = func(i int) {
		l, r, p := 2*i+1, 2*i+2, i
		if l < len(Q) && Q[l] > Q[p] {
			p = l
		}
		if r < len(Q) && Q[r] > Q[p] {
			r = p
		}
		if p != i {
			Q[i], Q[p] = Q[p], Q[i]
			Heapify(p)
		}
	}
	Pop := func() int {
		v := Q[0]
		Q[0], Q = Q[len(Q)-1], Q[:len(Q)-1]
		Heapify(0)
		return v
	}
	Push := func(v int) {
		Q = append(Q, v)
		i := len(Q) - 1
		for i > 0 && Q[i] > Q[(i-1)/2] {
			Q[i], Q[(i-1)/2] = Q[(i-1)/2], Q[i]
			i = (i - 1) / 2
		}
	}

	type Worker struct {
		r    float64
		w, q int
	}
	Workers := []Worker{}
	for i, q := range quality {
		Workers = append(Workers, Worker{float64(wage[i]) / float64(q), wage[i], q})
	}
	slices.SortFunc(Workers, func(a, b Worker) int { return a.w*b.q - a.q*b.w })
	log.Print(Workers)

	mCost, tQ := math.MaxFloat64, 0
	for _, worker := range Workers {
		log.Print(worker)

		tQ += worker.q

		Push(worker.q)
		log.Print(worker.q, " -> ", Q)

		if len(Q) > k {
			log.Print(" <- ", Q)
			tQ -= Pop()
		}

		if len(Q) == k {
			mCost = min(mCost, float64(tQ)*worker.r)
			log.Print(mCost, tQ, worker.q)
		}
	}
	return mCost
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 2331 Evaluate Boolean Binary Tree
func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val > 0
	}

	l, r := evaluateTree(root.Left), evaluateTree(root.Right)
	if root.Val == 2 {
		return l || r
	}
	return l && r
}

// 1325m Delete Leaves With a Given Value
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
