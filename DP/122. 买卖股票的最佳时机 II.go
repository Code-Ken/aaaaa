package DP

/*
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
            = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])

我们发现数组中的 k 已经不会改变了，也就是说不需要记录 k 这个状态了：
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
*/


func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp_i_0 := 0
	dp_i_1 := -prices[0]

	for i := 1; i < len(prices); i++ {
		tmp := dp_i_0
		dp_i_0 = maxDp1(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = maxDp1(dp_i_1, tmp-prices[i])
	}
	return dp_i_0
}

func maxDp1(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}