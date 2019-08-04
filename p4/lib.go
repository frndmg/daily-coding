package p4

// FindMissingMinPositiveV2 function
func FindMissingMinPositiveV2(l []int) int {
	n := len(l)

	get := func(i int) int {
		return l[i-1]
	}

	set := func(i, v int) { l[i-1] = v }

	swap := func(i, j int) int {
		l[i-1], l[j-1] = l[j-1], l[i-1]
		return get(i)
	}

	for i := 1; i <= n; i++ {
		g := get(i)
		for {
			if 1 <= g && g <= n && g != get(g) {
				g = swap(i, g)
				continue
			}

			break
		}

		if i != get(i) {
			set(i, 0)
		}
	}

	for i := 1; i <= n; i++ {
		if get(i) == 0 {
			return i
		}
	}

	return n + 1
}

// SolveExtraMemory function
func SolveExtraMemory(l []int) int {
	n := len(l)
	m := make([]bool, n+1)

	for i := 0; i < n; i++ {
		if 1 <= l[i] && l[i] <= n {
			m[l[i]] = true
		}
	}

	for i := 1; i <= n; i++ {
		if !m[i] {
			return i
		}
	}

	return n + 1
}

// // FindMissingMinPositive function
// func FindMissingMinPositive(v []int) int {
// 	if len(v) == 0 {
// 		return 1
// 	}
// 	return max(1, findMissingMinPositive(v, 1, 0, len(v)-1))
// }

// func findMissingMinPositive(v []int, x, lo, hi int) int {
// 	if lo >= hi {
// 		if v[hi] == x {
// 			return x + 1
// 		}
// 		return x
// 	}

// 	mid := (lo + hi) / 2

// 	x = findMissingMinPositive(v, x, lo, mid)
// 	x = findMissingMinPositive(v, x, mid+1, hi)

// 	return x
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
