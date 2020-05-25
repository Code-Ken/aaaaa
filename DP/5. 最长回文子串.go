package DP

func longestPalindrome(s string) string {
	if len(s) == 0{
		return ""
	}
	dp := make([][]int, len(s), len(s))
	for i, _ := range dp {
		dp[i] = make([]int, len(s), len(s))
	}

	maxL := 0
	ret := s[0:1]

	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] == 1 {
				if j-i+1 > maxL {
					maxL = j - i + 1
					ret = s[i : j+1]
				}
			}
		}
	}
	return ret

}