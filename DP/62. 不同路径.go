package DP

func uniquePaths(row int, col int) int {
	dp := make([][]int, row, row)
	for i, _ := range dp {
		dp[i] = make([]int, col, col)
	}

	//for (int i = 0; i < n; i++) dp[0][i]  = 1;//填充最左侧的
	//for (int j = 0; j < m; j++) dp[j][0] = 1;//填充最上侧的

	for i := 0; i < row; i++ {
		dp[i][col-1] = 1
	}

	for j := 0; j < col; j++ {
		dp[row-1][j] = 1
	}

	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}
