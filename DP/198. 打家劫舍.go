package DP

/*
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
*/

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	dp[1] = maxRob(nums[0], nums[1])
	ret := maxRob(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = maxRob(dp[i-1], dp[i-2]+nums[i])
		ret = maxRob(ret, dp[i])
	}
	return ret
}


func maxRob(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}