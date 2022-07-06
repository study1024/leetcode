package dushengchen

/**
Submission:
	https://leetcode.com/submissions/detail/739912367/

Runtime: 8 ms, faster than 53.25% of Go online submissions for Number of Islands.
Memory Usage: 3.9 MB, less than 57.88% of Go online submissions for Number of Islands.
 */
func numIslands(grid [][]byte) int {
	cnt := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			cnt++
			setLandToWater(grid, i, j)
		}
	}
	return cnt
}

var directions = [][]int{
	[]int{1, 0},
	[]int{0, 1},
	[]int{-1, 0},
	[]int{0, -1},
}
func setLandToWater(grid [][]byte, i, j int) {
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	for _, d := range directions {
		x, y := i+d[0], j+d[1]
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[i]) {
			continue
		}
		setLandToWater(grid, x, y)
	}
}