package dushengchen

import "fmt"

/*
Submission:
    https://leetcode.com/submissions/detail/362431123/
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	start := 0
	end := len(nums) - 1
	if nums[start] < nums[end] {
		pos, _ := binarySearch(nums, target)
		return pos
	}
	var mid int
	for {
		mid = (start + end) / 2
		fmt.Println("mid=", mid, "num[mid]=", nums[mid])
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[mid+1] {
			break
		}
		if nums[mid] > nums[0] {
			start = mid
		} else {
			end = mid
		}
	}
	fmt.Println("mid=", mid)
	if target > nums[mid] || target < nums[mid+1] {
		return -1
	}
	if target >= nums[0] {
		pos, _ := binarySearch(nums, target)
		return pos
	}
	x, _ := binarySearch(nums[mid+1:], target)
	if x == -1 {
		return -1
	}
	return mid + x + 1
}
