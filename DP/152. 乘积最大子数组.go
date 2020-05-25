package DP

/*
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

 

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。


*/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dpMax := make([]int, len(nums), len(nums))
	dpMin := make([]int, len(nums), len(nums))

	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	ret := nums[0]

	for i := 1; i < len(nums); i++ {
		dpMax[i] = maxP(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])
		dpMin[i] = minP(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])
		ret = maxP(dpMax[i], ret, ret)
	}
	return ret
}

func maxP(num1 int, num2 int, num3 int) int {
	if num1 >= num2 && num1 >= num3 {
		return num1
	}

	if num2 >= num1 && num2 >= num3 {
		return num2
	}
	return num3
}

func minP(num1 int, num2 int, num3 int) int {
	if num1 <= num2 && num1 <= num3 {
		return num1
	}
	if num2 <= num1 && num2 <= num3 {
		return num2
	}

	return num3
}
