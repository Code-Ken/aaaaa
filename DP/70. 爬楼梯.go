package DP

func climbStairs(n int) int {
	var dp = make([]int, n+1, n+1)
	steps := []int{1, 2}
	dp[0] = 1
	for j := 1; j < n+1; j++ {
		for i := 0; i < len(steps); i++ {
			if j-steps[i] < 0 {
				continue
			}
			dp[j] += dp[j-steps[i]]
		}
	}

	return dp[n]
}

func climbStairs(n int) int {
	f1, f2, f3 := 1, 2, 3
	if 1 == n {
		return f1
	}
	if 2 == n {
		return f2
	}
	for i := 2; i < n; i++ {
		f3 = f2 + f1
		f1 = f2
		f2 = f3
	}
	return f3
}
