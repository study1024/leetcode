package dushengchen

/**
Submission:
	https://leetcode.com/submissions/detail/372743143/

	Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations.
	Memory Usage: 2.7 MB, less than 51.11% of Go online submissions for Permutations.
	Next challenges:
		Permutations II
		Permutation Sequence
		Combinations
*/

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	ret := [][]int{}
	var x []int
	subs := permute(nums[1:])
	for i := 0; i < len(nums); i++ {
		for _, sub := range subs {
			x = InsertToSlice(sub, nums[0], i)
			// fmt.Println(i, nums, sub, x)
			ret = append(ret, x)
		}
	}
	return ret
}
