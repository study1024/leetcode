package dushengchen

/*
Submission:
   https://leetcode.com/submissions/detail/482436348/
*/

func exist(board [][]byte, word string) bool {
	byteWord := []byte(word)
	for i := range board {
		for j := range board[i] {
			if inner_exist(board, byteWord, i, j) {
				return true
			}
		}
	}
	return false
}
const space = ' '
func inner_exist(board [][]byte, word []byte, i, j int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] == space || board[i][j] != word[0]  {
		return false
	}
	if len(word) == 1 {
		return true
	}
	newWord := word[1:]
	var recoder byte
	recoder, board[i][j] = board[i][j], space
	if inner_exist(board, newWord, i-1, j) ||
		inner_exist(board, newWord, i+1, j) ||
		inner_exist(board, newWord, i, j-1) ||
		inner_exist(board, newWord, i, j+1) {
		return true
	}
	board[i][j] = recoder
	return false
}