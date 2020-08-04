package dushengchen

/**
Submission:
	https://leetcode.com/submissions/detail/375434494/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate Image.
Memory Usage: 2.2 MB, less than 68.18% of Go online submissions for Rotate Image.
*/
func rotate(matrix [][]int) {
	n := len(matrix) - 1
	for i := 0; i < n-i; i++ {
		for j := i; j < n-i; j++ {
			matrix[i][j], matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i] = matrix[n-j][i], matrix[i][j], matrix[j][n-i], matrix[n-i][n-j]
		}
	}
}
