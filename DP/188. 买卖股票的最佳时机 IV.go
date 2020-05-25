package DP

func maxProfit(maxK int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	//可以交易无限次
	if maxK > len(prices)/2 {
		dp_i_0 := 0
		dp_i_1 := -10000

		for i := 0; i < len(prices); i++ {
			tmp := dp_i_0
			dp_i_0 = maxDp3(dp_i_0, dp_i_1+prices[i])
			dp_i_1 = maxDp3(dp_i_1, tmp-prices[i])
		}
		return dp_i_0
	}

	dp := make([][][2]int, len(prices)+1, len(prices)+1)

	for i, _ := range dp {
		dp[i] = make([][2]int, maxK+1, maxK+1)
	}

	for i := 0; i < len(prices); i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = -10000
		for k := maxK; k >= 1; k-- {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = maxDp3(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = maxDp3(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[len(prices)-1][maxK][0]

}

func maxDp3(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
