package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/362744246/
*/

func searchRange(nums []int, target int) []int {
    var low int = 0
    var high int = len(nums) - 1
    for low <= high {
        mid := low + (high - low)/2
        if nums[mid] == target {
            start := mid
            end := mid
            for ; ; start-- {
                if start == 0 || nums[start-1] != target {
                    break
                }
            }
            for ; ; end++ {
                if end == (len(nums)-1) || nums[end+1] != target {
                    break
                }
            }
            return []int{start, end}
        } else if nums[mid] > target {
            high = mid -1
        } else {
            low = mid + 1
        }
    }
    return []int{-1, -1}
}