package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/359408757/
*/

func removeElement(nums []int, val int) int {
    i, j := -1, 0
    for {
        if j >= len(nums) {
            break
        }
        if nums[j] == val {
            j++
        } else {
            i++
            nums[i] = nums[j]
            j++
        }
    }
    return i+1
}