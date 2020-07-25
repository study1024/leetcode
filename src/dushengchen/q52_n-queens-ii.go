package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/369476316/
*/
func totalNQueens(n int) int {
	c := newCheckerboard(n)
	return solveNQueensRowN(c, 0)
}

func solveNQueensRowN(c checkerboard, row int) int {
	var res int
	for j, q := range c[row] {
		if q != 0 {
			continue
		}
		addQueen(c, row, j)
		// fmt.Println(row, j, c)
		if row < len(c[0])-1 {
			r := solveNQueensRowN(c, row+1)
			res = res + r
		} else {
			res = res + 1
		}
		delQueen(c, row, j)
	}
	// fmt.Println("res :=", res)
	return res
}
