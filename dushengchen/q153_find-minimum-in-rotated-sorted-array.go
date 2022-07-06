package dushengchen

/**
Submission:
   https://leetcode.com/submissions/detail/739924671/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Minimum in Rotated Sorted Array.
Memory Usage: 2.5 MB, less than 76.97% of Go online submissions for Find Minimum in Rotated Sorted Array.
*/
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[1]
		} else {
			return nums[0]
		}
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	start, end := 0,  len(nums) -1
	for {
		mid := int((start + end) / 2)
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		} else if nums[mid] < nums[mid-1]{
			return nums[mid]
		}
		if nums[mid] < nums[start] {
			end = mid
		} else {
			start = mid
		}
	}
}