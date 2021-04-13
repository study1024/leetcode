package main


/*
Submission:
    https://leetcode.com/submissions/detail/479940764/
*/

func setZeroes(matrix [][]int)  {
	firstRow, firstColumn := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColumn = true
			break
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			firstRow = true
			break
		}
	}

	for i:=1; i< len(matrix); i++ {
		for j:= 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] != 0 {
			continue
		}
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] = 0
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] != 0 {
			continue
		}
		for i := 1; i < len(matrix); i++ {
			matrix[i][j] = 0
		}
	}
	if firstRow {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if firstColumn {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{{1,2,3,4},{5,0,7,8},{0,10,11,12},{12,14,15,0}}
	setZeroes(matrix)
	println(matrix)
}