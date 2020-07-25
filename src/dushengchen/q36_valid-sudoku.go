package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/362749143/
*/


func isValidSudoku(board [][]byte) bool {

    for i := range board {
        rowMap := make(map[byte]bool, len(board[i]))
        colnmMap := make(map[byte]bool, len(board[i]))
        for j := range board[i] {
            if _, ok := rowMap[board[i][j]]; ok && board[i][j] != '.' {
                return false
            }
            if _, ok := colnmMap[board[j][i]]; ok && board[j][i] != '.' {
                return false
            }
            if (i == 0 || i == 3 || i == 6) && (j == 0 || j == 3 || j == 6) {
                arroudMap := make(map[byte]bool, len(board[i]))
                for x:= 0; x < 3; x++ {
                    for y :=0; y < 3; y++ {
                        if _, ok := arroudMap[board[i+x][j+y]]; ok && board[i+x][j+y] != '.' {
                            return false
                        }
                        arroudMap[board[i+x][j+y]] = true
                    }
                }
            }
            rowMap[board[i][j]], colnmMap[board[j][i]] = true, true
        }
    }
    return true
}