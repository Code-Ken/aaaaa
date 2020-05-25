package DP

import "math"

func isOneEditDistance(s string, t string) bool {
	lt := len(t)
	ls := len(s)
	if math.Abs(float64(lt-ls)) > 1.0 {
		return false
	}

	dp := make([][]int, lt+1, lt+1)
	for i, _ := range dp {
		dp[i] = make([]int, ls+1, ls+1)
	}
	for i := 0; i < lt+1; i++ {
		dp[i][0] = i
	}

	for j := 0; j < ls+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < lt+1; i++ {
		for j := 1; j < ls+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minI(
					dp[i-1][j-1]+1,
					dp[i-1][j]+1,
					dp[i][j-1]+1,
				)
			}
		}
	}

	return dp[lt][ls] == 1

}

func minI(num1 int, num2 int, num3 int) int {
	if num1 <= num2 && num1 <= num3 {
		return num1
	}
	if num2 <= num1 && num2 <= num3 {
		return num2
	}

	return num3
}
