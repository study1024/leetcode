package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/370068500/
*/
func maxProduct(nums []int) int {
    if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	dp := nums[0]
    dpm := nums[0]
	for i := 1; i < len(nums); i++ {
        if nums[i] >= 0 {
            dpm = dpm * nums[i]
            if nums[i] > dp * nums[i] {
                dp = nums[i]
            } else {
                dp = nums[i] * dp
            }
        } else {
            odp := dp
            dp = dpm * nums[i]
            if nums[i] < odp * nums[i] {
                dpm = nums[i]
            } else {
                dpm = odp * nums[i]
            }
        }
        if dp > max {
            max = dp
        }
        fmt.Println(i, nums[i], dp, dpm, max)
	}
	return max
}
