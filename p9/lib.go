package p9

// MaxSum function
func MaxSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := [2]int{}
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		if nums[i] < 0 || dp[0] < 0 {
			dp[0], dp[1] = dp[1], max(nums[i], dp[1])
		} else {
			dp[0], dp[1] = dp[1], max(dp[0]+nums[i], dp[1])
		}
	}

	return dp[1]
}

// MaxSumSlow function
func MaxSumSlow(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
		for j := 0; j < i-1; j++ {
			dp[i] = max(dp[i], dp[j]+nums[i])
		}
	}

	M := dp[0]
	for i := 1; i < len(dp); i++ {
		M = max(M, dp[i])
	}

	return M
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
