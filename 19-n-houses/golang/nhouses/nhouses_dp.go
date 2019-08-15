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

	// This part is not necesary
	// Colors are always different
	// ***

	// if x.k == a.k && x.c < a.c {
	// 	return x, b
	// }
	// if x.k == b.k && x.c < b.c {
	// 	return a, x
	// }

	// ***

	if x.c < a.c {
		return x, a
	} else if x.c < b.c {
		return a, x
	}

	return a, b
}
