package DP

func findLengthOfLCIS(nums []int) int {
	var dp = make([]int, len(nums), len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}

	//dp[i] = dp[i-1] + 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}
	ret := 0
	for _, val := range dp {
		if ret < val {
			ret = val
		}
	}
	return ret

}
