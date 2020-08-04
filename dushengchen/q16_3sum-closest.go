package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/282703014/
*/


func threeSumClosest(nums []int, target int) int {

    nums = SortInt(nums)
    var res int
    var min int
    min = -1
    for i, v := range nums {
        start, end := i+1, len(nums)-1
        if start == end {
            break
        }
        for {
            sum := v + nums[start]+nums[end]
            if target == sum {
                return sum
            }
            diff := AbsInt(target-sum)
            if diff < min || min == -1 {
                min = diff
                res = sum
            }
            if target < sum {
                end--
            } else {
                start++
            }
            if start == end {
                break
            }
        }
    }
    return res
}


