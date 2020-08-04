package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/370430163/
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	if obstacleGrid[0][0] != 0 {
		return 0
	}
	for i := range obstacleGrid {
		for j := range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i == 0 && j > 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
				continue
			}
			if j == 0 && i > 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
				continue
			}
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = -1
				continue
			}
			if obstacleGrid[i-1][j] < 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j] + obstacleGrid[i-1][j]
			}
			if obstacleGrid[i][j-1] < 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j] + obstacleGrid[i][j-1]
			}
		}
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] < 0 {
		return -obstacleGrid[m-1][n-1]
	}
	return 0
}
