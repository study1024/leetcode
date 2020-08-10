package dushengchen

import (
	"fmt"
	"strings"
)

/**
Submission:
	https://leetcode.com/submissions/detail/375848693/

Runtime: 12 ms, faster than 22.22% of Go online submissions for Valid Number.
Memory Usage: 2.6 MB, less than 66.67% of Go online submissions for Valid Number.
*/
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	hasE := -1
	hasPoint := -1
	hasM := -1
	hasNums := -1
	for i, c := range s {
		fmt.Println(string(c), hasE, hasM, hasPoint)
		switch c {
		case ' ':
			if hasE >= 0 || hasM >= 0 || hasNums >= 0 || hasM >= 0 {
				return false
			}
			continue
		case 'e':
			if hasE >= 0 || hasNums < 0 {
				return false
			}
			hasE = i
			hasPoint = i
			hasM = -1
			hasNums = -1
		case '.':
			if hasPoint >= 0 {
				return false
			}
			hasM = i
			hasPoint = i
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			hasM = i
			hasNums = i
		case '-', '+':
			if hasM >= 0 {
				return false
			}
			hasM = i
		default:
			return false
		}
	}
	if hasNums < 0 {
		return false
	}
	return true
}
