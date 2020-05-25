package DP

import "fmt"

/*
给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

 示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
*/

func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1, amount+1)
	for i, _ := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0

	//转移方程： dp[i] = min(dp[i-coin]+1,dp[i]) for coin in coins
	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if i-coin <0 {
				continue
			}
			if dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	fmt.Println(dp)
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}