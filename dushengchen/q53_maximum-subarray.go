package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/369640113/
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	dp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp+nums[i] {
			dp = nums[i]
		} else {
			dp = nums[i] + dp
		}
		if dp > max {
			max = dp
		}
	}
	return max
}
