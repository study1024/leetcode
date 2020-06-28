package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/359406846/
*/


func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    i, j := 0, 0
    for {
        if j >= len(nums) {
            break
        }
        if nums[i] == nums[j] {
            j++
        } else {
            i++
            nums[i] = nums[j]
            j++
        }
    }
    return i+1
}