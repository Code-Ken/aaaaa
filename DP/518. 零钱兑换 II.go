package DP

func change(amount int, coins []int) int {

	//dp[n] = ∑dp(n - step1)
	var dp = make([]int, amount+1, amount+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] < 0 {
				continue
			}
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}

