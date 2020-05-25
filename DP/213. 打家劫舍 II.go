package DP

//三个函数都用到了
func rob2(nums []int) int {

	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	dp[1] = maxRob2(nums[0], nums[1])
	ret := maxRob2(dp[0], dp[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = maxRob2(dp[i-1], dp[i-2]+nums[i])
		ret = maxRob2(ret, dp[i])
	}

	return ret
}

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) ==2 {
		return maxRob2(nums[0],nums[1])
	}

	return maxRob2(rob2(nums[0:len(nums)-1]), rob2(nums[1:]))
}

func maxRob2(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}