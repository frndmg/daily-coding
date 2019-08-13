package nhouses

import (
	"container/heap"
	"math"
)

type state struct {
	n, k, c int
}

type states []state

func (ss states) Less(i, j int) bool {
	return ss[i].c < ss[j].c
}

func (ss states) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss states) Len() int {
	return len(ss)
}

func (ss *states) Pop() interface{} {
	s := (*ss)[len(*ss)-1]
	*ss = (*ss)[:len(*ss)-1]
	return s
}

func (ss *states) Push(v interface{}) {
	s := v.(state)
	*ss = append(*ss, s)
}

// MinNHousesDisktra function
func MinNHousesDisktra(costs [][]int, N, K int) int {
	ss := make(states, 0)
	for k := 0; k < K; k++ {
		ss = append(ss, state{0, k, costs[0][k]})
	}
	heap.Init(&ss)

	m := math.MaxInt64

	for len(ss) > 0 {
		s := heap.Pop(&ss).(state)

		if s.n == N-1 {
			m = min(m, s.c)
			continue
		}

		for k := 0; k < K; k++ {
			if s.k == k {
				continue
			}

			heap.Push(&ss, state{s.n + 1, k, s.c + costs[s.n+1][k]})
		}
	}

	return m
}
