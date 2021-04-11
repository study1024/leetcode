package dushengchen

/*
Submission:
   https://leetcode.com/submissions/detail/479277903/
*/

func searchMatrix(matrix [][]int, target int) bool {
	high := len(matrix[0]) -1
	for i := range matrix {
		if matrix[i][high] == target {
			return true
		}
		if matrix[i][high] < target {
			continue
		}
		for j := range matrix[i] {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				return false
			}
		}
	}
	return false
}