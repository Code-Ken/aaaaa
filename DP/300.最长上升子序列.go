package DP

/**
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
 */

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums), len(nums))
	for k, _ := range dp {
		dp[k] = 1
	}

	//转移方程： dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)。
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
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