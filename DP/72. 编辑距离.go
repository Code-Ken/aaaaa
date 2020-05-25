package DP

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	dp := make([][]int, l1+1, l1+1)
	for i, _ := range dp {
		dp[i] = make([]int, l2+1, l2+1)
	}

	for i := 0; i < l1+1; i++ {
		dp[i][0] = i
	}

	for j := 0; j < l2+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minE(
					dp[i-1][j-1]+1,
					dp[i-1][j]+1,
					dp[i][j-1]+1)
			}
		}
	}
	return dp[l1][l2]

}

func minE(num1 int, num2 int, num3 int) int {
	if num1 <= num2 && num1 <= num3 {
		return num1
	}
	if num2 <= num1 && num2 <= num3 {
		return num2
	}

	return num3
}
