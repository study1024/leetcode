package dushengchen

/**
Submission:
	https://leetcode.com/submissions/detail/375851921/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
Memory Usage: 2 MB, less than 93.94% of Go online submissions for Plus One.
*/
func plusOne(digits []int) []int {
	before := 1
	for i := len(digits) - 1; i >= 0; i-- {
		before = before + digits[i]
		digits[i] = before % 10
		before = before / 10
		// fmt.Println(digits, i, before)
		if before == 0 {
			break
		}
	}
	if before == 0 {
		return digits
	}
	var x []int
	x = append(x, before)
	x = append(x, digits...)
	return x
}
