package DP

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	dp_i_0 := 0
	dp_i_1 := -prices[0] - fee

	for i := 1; i < len(prices); i++ {
		tmp := dp_i_0
		dp_i_0 = maxDp5(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = maxDp5(dp_i_1, tmp-prices[i]-fee)
	}
	return dp_i_0
}

func maxDp5(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
