package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/370152730/
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	is, ie, js, je := 0, len(matrix)-1, 0, len(matrix[0])-1
	istep, jstep := 0, 1
	i, j := is, js
	ret := make([]int, len(matrix[0])*len(matrix))
	idx := 0
	for idx < cap(ret) {
		// fmt.Println(i, j, istep, jstep, matrix[i][j])
		ret[idx] = matrix[i][j]
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
	return ret
}
