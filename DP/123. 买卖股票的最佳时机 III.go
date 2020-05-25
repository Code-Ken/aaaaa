package DP

/*
dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
dp[i][1][1] = max(dp[i-1][1][1], -prices[i])
    int dp_i10 = 0, dp_i11 = Integer.MIN_VALUE;
    int dp_i20 = 0, dp_i21 = Integer.MIN_VALUE;

任意k次
int maxProfit_k_any(int max_k, int[] prices) {
    int n = prices.length;
    if (max_k > n / 2)
        return maxProfit_k_inf(prices);

    int[][][] dp = new int[n][max_k + 1][2];
    for (int i = 0; i < n; i++)
        for (int k = max_k; k >= 1; k--) {
            if (i - 1 == -1) {  处理 base case }
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i]);
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i]);
		}
	return dp[n - 1][max_k][0];
}


*/

const Inf = -100000

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp_i20 := 0
	dp_i10 := 0
	dp_i11 := Inf
	dp_i21 := Inf

	for i := 0; i < len(prices); i++ {
		dp_i20 = maxDp2(dp_i20, dp_i21+prices[i])
		dp_i21 = maxDp2(dp_i21, dp_i10-prices[i])
		dp_i10 = maxDp2(dp_i10, dp_i11+prices[i])
		dp_i11 = maxDp2(dp_i11, - prices[i])
	}
	return dp_i20
}

func maxDp2(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}