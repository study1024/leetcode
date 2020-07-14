package dushengchen

import "fmt"

/**
Submission:
    https://leetcode.com/submissions/detail/366393972/
*/
func firstMissingPositive(nums []int) int {
    for i, v := range  nums {
        if v <= 0 {
            nums[i] = -2
        }
    }
    for _, v := range  nums {
        set(nums, v)
    }
    fmt.Println(nums)
    for i, v := range  nums {
        if v != -1 {
            return i+1
        }
    }
    return len(nums)+1
}

func set(nums []int, target int) {
    if target > len(nums) || target <= 0 {
        return
    }
    x := nums[target-1]
    nums[target-1] = -1
    if x != target {
        set(nums, x)
    }
}