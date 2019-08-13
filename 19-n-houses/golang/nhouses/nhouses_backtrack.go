package nhouses

import "math"

// MinNHousesBacktrack function
func MinNHousesBacktrack(costs [][]int, N, K int) int {
	return minNHousesRec(costs, N, K, 0, 0, -1, math.MaxInt64)
}

func minNHousesRec(costs [][]int, N, K, i, s, p, m int) int {
	if s >= m {
		return m
	}

	if i >= N {
		return s
	}

	for k := 0; k < K; k++ {
		if k == p { // same color
			continue
		}

		c := costs[i][k]
		m = min(m, minNHousesRec(costs, N, K, i+1, s+c, k, m))
	}

	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
