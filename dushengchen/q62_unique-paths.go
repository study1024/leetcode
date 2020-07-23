package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/370424455/
*/

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	x := NewIntArray2D(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				x[i][j] = 1
				continue
			}
			x[i][j] = x[i-1][j] + x[i][j-1]
		}
	}
	return x[m-1][n-1]
}

//time limit
// func uniquePaths(m int, n int) int {
// 	x := NewIntArray2D(m, n)
// 	return uniquePathsAtPos(x, 0, 0)
// }
// func uniquePathsAtPos(b [][]int, i, j int)  int {
// 	my := 0
// 	if i < 0 || i>= len(b) {
// 		return 0
// 	}
// 	if j < 0 || j>= len(b[i]) {
// 		return 0
// 	}
// 	if b[i][j] != 0 {
// 		return 0
// 	}
// 	if i == (len(b)-1) && (j == len(b[i])-1) {
// 		return 1
// 	}
// 	b[i][j] = 1
// 	my = my + uniquePathsAtPos(b, i+1, j)
// 	my = my + uniquePathsAtPos(b, i, j+1)
//     b[i][j] = 0
// 	return my
// }
