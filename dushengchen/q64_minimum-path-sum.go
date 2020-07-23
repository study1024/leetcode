package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/370432851/
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := range grid {
		for j := range grid[i] {
			if i == 0 && j > 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
				continue
			}
			if j == 0 && i > 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
				continue
			}
			if i == 0 && j == 0 {
				continue
			}
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			} else {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			}
		}
	}
	m, n := len(grid), len(grid[0])
	return grid[m-1][n-1]
}
