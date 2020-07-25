package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/370577886/
*/
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	reachable := make([]bool, len(nums))
	reachable[0] = true
	for i, v := range nums {
		if !reachable[i] {
			continue
		}
		for j := 1; j <= v; j++ {
			if (i + j) >= (len(nums) - 1) {
				return true
			}
			reachable[i+j] = true
		}
	}
	return reachable[len(reachable)-1]
}
