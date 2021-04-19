package dushengchen


/*
Submission:
	https://leetcode.com/submissions/detail/482442786/
*/

func removeDuplicates2(nums []int) int {
	return removeDuplicatesN(nums, 2)
}

func removeDuplicatesN(nums []int, n int) int {
	if n < 1 {
		return 0
	}
	l := 0
	d := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[l] {
			l, d = l+1, 1
			nums[l] = nums[i]
			continue
		}
		if d < n {
			l, d = l+1, d+1
			nums[l] = nums[i]
			continue
		}
	}
	return l+1
}