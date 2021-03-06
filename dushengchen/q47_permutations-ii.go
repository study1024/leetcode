package dushengchen

/**
Submission:
	https://leetcode.com/submissions/detail/372752824/

Runtime: 4 ms, faster than 77.21% of Go online submissions for Permutations II.
Memory Usage: 3.5 MB, less than 77.78% of Go online submissions for Permutations II.
*/

func permuteUnique(nums []int) [][]int {
	nums = SortInt(nums)
	var ret [][]int
	ret = append(ret, CopyInt(nums))
	for {
		if x := nextPermutationX(nums); !x {
			break
		}
		ret = append(ret, CopyInt(nums))
	}
	return ret
}

//这段代码来源q31，不同的是，源代码会在完全逆序时翻转，而本场景不需要，所以返回fasle
func nextPermutationX(nums []int) bool {
	if len(nums) <= 1 {
		return false
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
		return false
	}
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
	RevertInt(nums[i:])
	return true
}
