package dushengchen

/**
Submission:
   https://leetcode.com/submissions/detail/369642600/
*/
func maxProfit(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	max := 0
	dp := 0
	for i := 1; i < len(nums); i++ {
		if (nums[i] - nums[i-1]) > (nums[i] - nums[i-1] + dp) {
			dp = nums[i] - nums[i-1]
		} else {
			dp = nums[i] - nums[i-1] + dp
		}
		if dp > max {
			max = dp
		}
	}
	return max
}
