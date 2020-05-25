package DP

func minimumTotal(triangle [][]int) int {

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			triangle[i][j] += minM(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]

}

func minM(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

