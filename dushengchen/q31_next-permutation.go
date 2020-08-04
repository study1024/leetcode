package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/371331223/
*/

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var i int
	find := false
	for i = len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			find = true
			break
		}
	}
	if !find {
		RevertInt(nums)
		return
	}
	// fmt.Println(i, nums[i])
	var j int
	for j = i; true; j++ {
		if j+1 == len(nums) {
			nums[i-1], nums[j] = nums[j], nums[i-1]
			break
		}
		if nums[i-1] < nums[j] && nums[i-1] >= nums[j+1] {
			nums[i-1], nums[j] = nums[j], nums[i-1]
			break
		}
	}
	// fmt.Println(j, nums[j], nums)
	RevertInt(nums[i:])
	return
}
