package DP


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])

	dp := make([][]int, row, row)
	for i, _ := range dp {
		dp[i] = make([]int, col, col)
	}
	for i := row - 1; i >= 0; i-- {
		if obstacleGrid[i][col-1] == 1 {
			break
		}
		dp[i][col-1] = 1
	}


	for j := col - 1; j >= 0; j-- {
		if obstacleGrid[row-1][j] == 1 {
			break
		}
		dp[row-1][j] = 1
	}


	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}