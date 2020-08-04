package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/371274059/
*/
func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	matrix := NewIntArray2D(n, n)

	is, ie, js, je := 0, len(matrix)-1, 0, len(matrix[0])-1
	istep, jstep := 0, 1
	i, j := is, js
	idx := 1
	for idx <= n*n {
		// fmt.Println(i, j, istep, jstep, matrix[i][j])
		matrix[i][j] = idx
		idx++
		i, j = i+istep, j+jstep
		if j > je {
			i, j, istep, jstep, is = is+1, je, 1, 0, is+1
		} else if i > ie {
			i, j, istep, jstep, je = ie, j-1, 0, -1, je-1
		} else if j < js {
			i, j, istep, jstep, ie = ie-1, js, -1, 0, ie-1
		} else if i < is {
			i, j, istep, jstep, js = is, j+1, 0, 1, js+1
		}
	}
	return matrix
}
