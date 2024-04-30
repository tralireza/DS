package main

import (
	"log"
	"testing"
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
