package p1

func twosumk(list []int, k int) bool {
	m := make(map[int]struct{})

	for _, v := range list {
		if _, ok := m[k-v]; ok {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}
