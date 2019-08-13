package nhouses

import "math"

type dpState struct {
	c, k int
}

// MinNHousesDP function
func MinNHousesDP(costs [][]int, N, K int) int {
	x, y := dpState{math.MaxInt64, -1}, dpState{math.MaxInt64, -2}

	for k := 0; k < K; k++ {
		x, y = update(x, y, dpState{costs[0][k], k})
	}

	for i := 1; i < N; i++ {
		a, b := dpState{math.MaxInt64, -1}, dpState{math.MaxInt64, -2}
		for k := 0; k < K; k++ {
			z := x
			if z.k == k {
				z = y
			}

			a, b = update(a, b, dpState{z.c + costs[i][k], k})
		}
		x, y = a, b
	}

	return x.c
}

func update(a, b, x dpState) (dpState, dpState) {
	// a.k != b.k
	// a.c <= b.c
	if x.c < a.c {
		if x.k != a.k {
			return x, a
		}
		return x, b
	} else if x.c < b.c {
		if x.k != a.k {
			return a, x
		}
		return x, b
	}
	return a, b
}
