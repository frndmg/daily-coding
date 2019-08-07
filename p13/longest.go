package p13

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LongestSubstringSlow function
func LongestSubstringSlow(k int, s string) int {
	count := func(i int, j int) int {
		m := make(map[rune]struct{})
		for ; i <= j; i++ {
			m[rune(s[i])] = struct{}{}
		}
		return len(m)
	}

	maxLength := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if count(i, j) <= k {
				maxLength = max(maxLength, j-i+1)
			}
		}
	}

	return maxLength
}

// LongestSubstring function
func LongestSubstring(k int, s string) int {
	if k <= 0 {
		return 0
	}

	m := make(map[rune]int)

	i := 0
	j := 0

	maxLength := 0

	for ; j < len(s); j++ {
		m[rune(s[j])]++

		for ; len(m) > k && i < j; i++ {
			r := rune(s[i])

			m[r]--
			if m[r] == 0 {
				delete(m, r)
			}
		}

		maxLength = max(maxLength, j-i+1)
	}

	return maxLength
}
