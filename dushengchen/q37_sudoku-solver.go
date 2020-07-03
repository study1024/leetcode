package dushengchen


/**
Submission:
    https://leetcode.com/submissions/detail/361497868/

test input:
    [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
or
    [[".",".",".","2",".",".",".","6","3"],["3",".",".",".",".","5","4",".","1"],[".",".","1",".",".","3","9","8","."],[".",".",".",".",".",".",".","9","."],[".",".",".","5","3","8",".",".","."],[".","3",".",".",".",".",".",".","."],[".","2","6","3",".",".","5",".","."],["5",".","3","7",".",".",".",".","8"],["4","7",".",".",".","1",".",".","."]]
 */
func solveSudoku(board [][]byte)  {
    a := covertToCells(board)
    if loopSearch(a) {
        for i := range a{
            for j := range a[i] {
                board[i][j] = a[i][j].Fill
            }
        }
    }
    return
}

func loopSearch(a [][]*Cell) bool {
    minAllowCellCnt := 10
    mX, mY := -1, -1
    getAnswer := true
    block := false
    for i := range a{
        for j := range a[i] {
            if a[i][j].Fill == '.' {
                getAnswer = false
            }
            if a[i][j].Fill == '.' && len(a[i][j].Allow) == 0 {
                block = true
                break
            }
            if a[i][j].Fill == '.' && len(a[i][j].Allow) < minAllowCellCnt {
                mX, mY, minAllowCellCnt = i, j, len(a[i][j].Allow)
            }
        }
    }
    if block {
        return false
    }
    if getAnswer {
        return true
    }

    for f, _ := range a[mX][mY].Allow {
        if !preFillCell(a, f, mX, mY) {
            continue
        }
        a[mX][mY].Fill = f
        drops := dropCellAllows(a, f, mX, mY)
        if loopSearch(a) {
            return true
        }
        //恢复现场
        a[mX][mY].Fill = '.'
        addCellAllows(a, f, drops)
    }
    return false
}


type Cell struct {
    Fill byte
    Allow   map[byte]bool
}

var r = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
func newAllowMap()  map[byte]bool {
    m := make( map[byte]bool, 9)
    for _, x:= range r {
        m[x] = true
    }
    return m
}

func covertToCells(board [][]byte)([][]*Cell) {
    a := [][]*Cell{}
    for i := range board {
        row := make([]*Cell, len(board[i]))
        for j := range row {
            row[j] = &Cell{
                Fill: '.',
                Allow: newAllowMap(),
            }
        }
        a = append(a, row)
    }
    for i := range board {
        for j := range board[i] {
            if board[i][j] != '.' {
                a[i][j].Fill = board[i][j]
                dropCellAllows(a, board[i][j], i, j)
            }
        }
    }
    return a
}

//func fillCell(a [][]*Cell, fill byte, i, j int) [][]int {
//    //已经填充了
//    a[i][j].Fill = fill
//    return dropCellAllows(a, fill, i, j)
//}

//准备填充前的检查
func preFillCell(a [][]*Cell, fill byte, i, j int) bool {
    //已经填充了
    if a[i][j] == nil || a[i][j].Fill != '.' {
        return false
    }
    for x := 0; x < 9; x++ {
        if a[i][x].Fill == fill || a[x][j].Fill == fill {
            return false
        }
    }
    xStart := i - (i % 3)
    yStart := j - (j % 3)
    for x:= 0; x < 3; x++ {
        for y :=0; y < 3; y++ {
            if a[xStart+x][yStart+y].Fill == fill {
                return false
            }
        }
    }
    return true
}

//drop相关cell的allow中的字符，并记录下drop的位置
func dropCellAllows(a [][]*Cell, drop byte, i, j int) [][]int {
    ret := [][]int{}
    for x := 0; x < 9; x++ {
        if a[i][x].Allow[drop] {
            delete(a[i][x].Allow, drop)
            ret = append(ret, []int{i, x})
        }
        if a[x][j].Allow[drop] {
            delete(a[x][j].Allow, drop)
            ret = append(ret, []int{x, j})
        }
    }
    xStart := i - (i % 3)
    yStart := j - (j % 3)
    for x:= 0; x < 3; x++ {
        for y :=0; y < 3; y++ {
            if a[xStart+x][yStart+y].Allow[drop] {
                delete(a[xStart+x][yStart+y].Allow, drop)
                ret = append(ret, []int{xStart+x, yStart+y})
            }
        }
    }
    return ret
}

func addCellAllows(a [][]*Cell, add byte, pos [][]int) {
    for _, item := range pos {
        a[item[0]][item[1]].Allow[add] = true
    }
}