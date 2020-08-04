package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/369470096/
*/
func solveNQueens(n int) [][]string {
	c := newCheckerboard(n)
	return solveNQueensRow(c, 0)
}

type checkerboard [][]int

func newCheckerboard(n int) checkerboard {
	x := [][]int{}
	for i := 0; i < n; i++ {
		x = append(x, make([]int, n))
	}
	return checkerboard(x)
}

func solveNQueensRow(c checkerboard, row int) [][]string {
	res := [][]string{}
	for j, q := range c[row] {
		if q != 0 {
			continue
		}
		addQueen(c, row, j)
		// fmt.Println(row, j, c)
		if row < len(c[0])-1 {
			r := solveNQueensRow(c, row+1)
			if len(r) > 0 {
				res = append(res, r...)
			}
		} else {
			res = append(res, c.Covert())
		}
		delQueen(c, row, j)
	}
	// fmt.Println("res :=", res)
	return res
}

func addQueen(board checkerboard, i, j int) {
	//todo
	interval1 := j - i
	interval2 := j + i

	for x := 0; x < len(board[0]); x++ {
		board[i][x] = board[i][x] - 1
		board[x][j] = board[x][j] - 1
		j1 := x + interval1
		j2 := interval2 - x
		if j1 >= 0 && j1 < len(board[0]) {
			board[x][j1] = board[x][j1] - 1
		}
		if j2 >= 0 && j2 < len(board[0]) {
			board[x][j2] = board[x][j2] - 1
		}
	}
	board[i][j] = 1
}

func delQueen(board checkerboard, i, j int) {
	interval1 := j - i
	interval2 := j + i

	for x := 0; x < len(board[0]); x++ {
		board[i][x] = board[i][x] + 1
		board[x][j] = board[x][j] + 1
		j1 := x + interval1
		j2 := interval2 - x
		if j1 >= 0 && j1 < len(board[0]) {
			board[x][j1] = board[x][j1] + 1
		}
		if j2 >= 0 && j2 < len(board[0]) {
			board[x][j2] = board[x][j2] + 1
		}
	}
	board[i][j] = 0
}

func (b checkerboard) Covert() []string {
	ret := make([]string, len(b[0]))
	for i := range b {
		x := make([]rune, len(b[0]))
		for j := range b[i] {
			if b[i][j] == 1 {
				x[j] = 'Q'
			} else {
				x[j] = '.'
			}
		}
		ret[i] = string(x)
	}
	return ret
}
