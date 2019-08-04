package p7

// Count function
func Count(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	dp := [2]int{}

	twoCharsFitAt := func(idx int) bool {
		return s[idx-1] == '1' || (s[idx-1] == '2' && s[idx] <= '6')
	}

	dp[0] = 1
	if twoCharsFitAt(1) {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	for i := 2; i < len(s); i++ {
		// make
		// an = an-1 + an-2
		// or
		// an = an-1
		if twoCharsFitAt(i) {
			dp[0], dp[1] = dp[1], dp[1]+dp[0]
		} else {
			dp[0] /*, dp[1] */ = dp[1] /*, dp[1] */
		}
	}

	return dp[1]
}
