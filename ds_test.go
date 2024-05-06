package DS

import (
	"container/list"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")

	log.Print("> DS")
}

// 834h Sum of Distances in Tree
func Test834(t *testing.T) {
	sumOfDistancesInTree := func(n int, edges [][]int) []int {
		lsAdj := make([][]int, n)

		for _, edge := range edges {
			v, u := edge[0], edge[1]
			lsAdj[v] = append(lsAdj[v], u)
			lsAdj[u] = append(lsAdj[u], v)
		}

		dists, subTreeNds := make([]int, n), make([]int, n)

		var walk func(v, p int)
		walk = func(v, p int) {
			subTreeNds[v] = 1
			for _, u := range lsAdj[v] {
				if u != p {
					walk(u, v)
					subTreeNds[v] += subTreeNds[u]
					dists[v] += dists[u] + subTreeNds[u]

					log.Print(v, u, dists)
				}
			}
		}
		walk(0, -1)
		log.Print(dists, subTreeNds)

		var reRoot func(v, p int)
		reRoot = func(v, p int) {
			for _, u := range lsAdj[v] {
				if u != p {
					dists[u] = (dists[v] - subTreeNds[u]) + (n - subTreeNds[u])
					reRoot(u, v)
				}
			}
		}
		reRoot(0, -1)
		log.Print(dists, subTreeNds)

		return dists
	}

	log.Print("[8 12 6 10 10 10] ?= ", sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}))
	log.Print("[1 1] ?= ", sumOfDistancesInTree(2, [][]int{{0, 1}}))
}

// 2997m Minimum Number of Operations to Make Array XOR Equal to K
func Test2997(t *testing.T) {
	minOperations := func(nums []int, k int) int {
		x := k
		for _, n := range nums {
			x ^= n
		}

		ops := 0
		for x > 0 {
			ops += x & 1
			x >>= 1
		}
		return ops
	}

	log.Print("2 ?= ", minOperations([]int{2, 1, 3, 4}, 1))
	log.Print("0 ?= ", minOperations([]int{2, 0, 2, 0}, 0))
}

// 1915m Number of Wonderful Substrings
func Test1915(t *testing.T) {
	wonderfulSubstrings := func(word string) int64 {
		var x int64

		Frq := map[int]int64{}
		Frq[0] = 1

		mask := 0
		for i := 0; i < len(word); i++ {
			mask ^= 1 << (word[i] - 'a')

			log.Printf("%3d %010b   %q", mask, mask&0x3ff, word[0:i+1])

			if f, ok := Frq[mask]; ok {
				x += f
				Frq[mask]++
			} else {
				Frq[mask] = 1
			}

			for p := 0; p < 10; p++ {
				if f, ok := Frq[mask^(1<<p)]; ok {
					x += f
				}
			}
		}

		log.Print(Frq)

		return x
	}

	log.Print("3 ?= ", wonderfulSubstrings("ab"))
	log.Print("4 ?= ", wonderfulSubstrings("aab"))
	log.Print("9 ?= ", wonderfulSubstrings("aabb"))
	log.Print("12 ?= ", wonderfulSubstrings("feffaec"))
}

// 2000 Reverse Prefix of Word
func Test2000(t *testing.T) {
	reversePrefix := func(word string, ch byte) string {
		bs := []byte(word)
		i := 0
		for i < len(bs) && word[i] != ch {
			i++
		}
		if i < len(bs) {
			j := 0
			for j < i {
				bs[j], bs[i] = bs[i], bs[j]
				i--
				j++
			}
		}
		return string(bs)
	}

	stacker := func(word string, ch byte) string {
		S := list.New()

		bs := make([]byte, 0, len(word))
		for i := 0; i < len(word); i++ {
			if word[i] == ch {
				bs = append(bs, word[i])
				for S.Len() > 0 {
					bs = append(bs, S.Remove(S.Back()).(byte))
				}
				return string(bs) + word[i+1:]
			}
			S.PushBack(word[i])
		}

		return word
	}

	for _, f := range []func(string, byte) string{reversePrefix, stacker} {
		log.Print("dcbaefd ?= ", f("abcdefd", 'd'))
		log.Print("zxyxxe ?= ", f("xyxzxe", 'z'))
		log.Print("abcd ?= ", f("abcd", 'z'))
	}
}

// 2441 Largest Positive Integer That Exists With Its Negative
func Test2441(t *testing.T) {
	// ~O(n)
	findMaxK := func(nums []int) int {
		x := -1
		Mem := map[int]struct{}{}
		for _, n := range nums {
			if _, ok := Mem[-n]; ok {
				x = max(x, max(n, -n))
			} else {
				Mem[n] = struct{}{}
			}
		}
		return x
	}

	// O(nlogn)
	twoPointers := func(nums []int) int {
		slices.Sort(nums)
		l, r := 0, len(nums)-1
		for l < r {
			if -nums[l] == nums[r] {
				return nums[r]
			} else if -nums[l] < nums[r] {
				r--
			} else {
				l++
			}
		}
		return -1
	}

	// O(n)
	// 1 <= n <= 100_000
	hashArray := func(nums []int) int {
		Mem := make([]int, 100_001)

		x := -1
		for _, n := range nums {
			if n < 0 {
				if Mem[-n] == -n {
					x = max(x, -n)
				} else {
					Mem[-n] = n
				}
			} else {
				if Mem[n] == -n {
					x = max(x, n)
				} else {
					Mem[n] = n
				}
			}
		}
		return x
	}

	nums := make([]int, 0, 100_000)
	for range 100_000 {
		n := rand.Intn(100_000) + 1
		if rand.Intn(2) == 1 {
			n *= -1
		}
		nums = append(nums, n)
	}
	for _, f := range []func([]int) int{findMaxK, twoPointers, hashArray} {
		log.Print("3 ?= ", f([]int{-1, 2, -3, 3}))
		log.Print("7 ?= ", f([]int{-1, 10, 6, 7, -7, 1}))
		log.Print("-1 ?= ", f([]int{-10, 8, 6, 7, -2, -3}))
		ts := time.Now()
		v := f(nums)
		dur := time.Since(ts)
		log.Print(" ?= ", v, " [", dur, "]")
		log.Print("===")
	}
}

// 165m Compare Version Numbers
func Test165(t *testing.T) {
	compareVersion := func(version1 string, version2 string) int {
		Ver1, Ver2 := strings.Split(version1, "."), strings.Split(version2, ".")
		for i1, i2 := 0, 0; i1 < len(Ver1) || i2 < len(Ver2); i1, i2 = i1+1, i2+1 {
			v1, v2 := 0, 0
			if i1 < len(Ver1) {
				v1, _ = strconv.Atoi(Ver1[i1])
			}
			if i2 < len(Ver2) {
				v2, _ = strconv.Atoi(Ver2[i2])
			}

			if v1 < v2 {
				return -1
			} else if v1 > v2 {
				return 1
			}
		}
		return 0
	}

	log.Print("0 ?= ", compareVersion("1.01", "1.001"))
	log.Print("0 ?= ", compareVersion("1.0", "1.0.0"))
	log.Print("-1 ?= ", compareVersion("0.1", "1.1"))
}

// 881m Boats to Save People
func Test881(t *testing.T) {
	numRescueBoats := func(people []int, limit int) int {
		slices.Sort(people)

		boats := 0
		l, r := 0, len(people)-1
		for l <= r {
			boats++
			if people[l]+people[r] <= limit {
				l++
			}
			r--
		}
		return boats
	}

	log.Print("1 ?= ", numRescueBoats([]int{1, 2}, 3))
	log.Print("3 ?= ", numRescueBoats([]int{3, 2, 2, 1}, 3))
	log.Print("4 ?= ", numRescueBoats([]int{3, 5, 3, 4}, 5))
}

// 237m Delete Node in a Linked List
func Test237(t *testing.T) {
	type ListNode struct {
		Val  int
		Next *ListNode
	}

	// node Is an Internal Node
	deleteNode := func(node *ListNode) {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}

	draw := func(n *ListNode) {
		for ; n != nil; n = n.Next {
			if n.Next != nil {
				fmt.Printf("{%d *}->", n.Val)
			} else {
				fmt.Printf("{%d /}\n", n.Val)
			}
		}
	}

	type L = ListNode
	n := &L{5, &L{1, &L{Val: 9}}}
	l := &L{4, n}

	draw(l)
	deleteNode(n)
	draw(l)
}

// 2487m Remove Nodes from Linked List
func Test2487(t *testing.T) {
	withStack := func(head *ListNode) *ListNode {
		S := list.New()
		for n := head; n != nil; n = n.Next {
			for S.Len() > 0 && n.Val > S.Back().Value.(*ListNode).Val {
				S.Remove(S.Back())
			}
			S.PushBack(n)
		}

		head = S.Remove(S.Front()).(*ListNode)
		n := head
		for S.Len() > 0 {
			n.Next = S.Remove(S.Front()).(*ListNode)
			n = n.Next
		}
		n.Next = nil
		return head
	}

	draw := func(n *ListNode) {
		for ; n != nil; n = n.Next {
			if n.Next != nil {
				fmt.Printf("{%d *}->", n.Val)
			} else {
				fmt.Printf("{%d /}\n", n.Val)
			}
		}
	}

	type L = ListNode
	for _, l := range []*L{&L{5, &L{2, &L{13, &L{3, &L{Val: 8}}}}}, &L{1, &L{1, &L{1, &L{Val: 1}}}}} {
		for _, f := range []func(*ListNode) *ListNode{withStack, removeNodes} {
			draw(l)
			draw(f(l))
		}
		log.Print("+++")
	}
}
