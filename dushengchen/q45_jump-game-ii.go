package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/370580391/
*/
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	reachStep := make([]int, len(nums))
	reachStep[0] = 0
	for i, v := range nums {
		if i > 0 && reachStep[i] <= 0 {
			continue
		}
		for j := 1; j <= v; j++ {
			if (i + j) > (len(nums) - 1) {
				break
			}
			if reachStep[i+j] == 0 {
				reachStep[i+j] = reachStep[i] + 1
			} else if reachStep[i+j] > (reachStep[i] + 1) {
				reachStep[i+j] = reachStep[i] + 1
			}
		}
	}
	return reachStep[len(reachStep)-1]
}
