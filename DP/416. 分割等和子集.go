package DP

/*
如果不把 nums[i] 算入子集，或者说你不把这第 i 个物品装入背包
那么是否能够恰好装满背包，取决于上一个状态 dp[i-1][j]，继承之前的结果。

如果把 nums[i] 算入子集，或者说你把这第 i 个物品装入了背包
那么是否能够恰好装满背包，取决于状态 dp[i - 1][j-nums[i-1]]。

*/

func canPartition2(nums []int) bool {
	sum := sumSlice(nums)
	if sum%2 == 1 {
		return false
	}

	total := sum / 2
	dp := make([]bool,total+1)
	dp[0] = true
	for i:=0;i<len(nums);i++{
		for j:=total;j>=0;j--{
			if j-nums[i] >=0{
				dp[j] = dp[j] || dp[j-nums[j]]
			}
		}
	}
	return dp[total]
}


func canPartition(nums []int) bool {
	sum := sumSlice(nums)
	if sum%2 == 1 {
		return false
	}

	total := sum / 2
	dp := make([][]bool, len(nums)+1, len(nums)+1)

	for i, _ := range dp {
		dp[i] = make([]bool, total+1, total+1)
		dp[i][0] = true
	}

	for i := 1; i < len(nums)+1; i++ {
		for j := 1; j < total+1; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[len(nums)][total]
}


func sumSlice(nums []int) int {
	ret := 0
	for _, val := range nums {
		ret += val
	}
	return ret
}

