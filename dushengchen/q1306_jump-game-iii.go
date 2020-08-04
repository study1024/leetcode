package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/370587342/
*/

func canReach(arr []int, start int) bool {
	if len(arr) == 0 {
		return true
	}
	reachStep := make([]int, len(arr))
	return canReachRv(arr, reachStep, start)
}

func canReachRv(nums []int, reachStep []int, start int) bool {
	if start < 0 || start >= len(reachStep) {
		return false
	}
	if reachStep[start] > 0 {
		return false
	}
	reachStep[start] = 1
	if nums[start] == 0 {
		return true
	}
	if canReachRv(nums, reachStep, start+nums[start]) {
		return true
	}
	if canReachRv(nums, reachStep, start-nums[start]) {
		return true
	}
	return false
}
