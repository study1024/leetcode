package dushengchen

import "strings"

/*
question:
    https://leetcode.com/problems/zigzag-conversion/
Submission:
    https://leetcode.com/submissions/detail/289227013/

*/

func convert(s string, numRows int) string {
    if numRows <= 1 {
        return s
    }
    rows := make([]string, numRows)
    idx := 0
    step:= 1
    for _, r:= range s {
        rows[idx] = string(append([]rune(rows[idx]), r))
        idx = idx + step
        if idx == numRows {
            idx = numRows - 2
            step = -1
        } else if idx < 0 {
            idx = 1
            step = 1
        }
    }
    return strings.Join(rows, "")
}