package dushengchen

import "fmt"

/*
Submission:
    https://leetcode.com/submissions/detail/356983445/
*/

func fourSum(nums []int, target int) [][]int {
    //nums = SortInt(nums)

    numsMap := map[int][][]int{}
    var res [][]int
    for i:=0; i<len(nums); i++ {
        for j:=i+1; j<len(nums); j++ {
            numsMap[nums[i]+nums[j]] = append(numsMap[nums[i]+nums[j]], []int{i, j})
        }
    }
    resMap := map[string]bool{}
    for midNum, arr1 := range numsMap {
        other := target - midNum
        if other > midNum {
            continue
        }
        if arr2, ok  := numsMap[other]; ok && len(arr2) > 0 {
            for i,item1 := range arr1 {
                for j,item2 := range arr2 {
                    if other == midNum && i<j {
                        continue
                    }
                    if item1[0] == item2[0] || item1[1] == item2[1] ||
                        item1[0] == item2[1] || item1[1] == item2[0] {
                        continue
                    }
                    h := HashInt(nums[item1[0]], nums[item1[1]], nums[item2[0]], nums[item2[1]])
                    if resMap[h] {
                        continue
                    } else {
                        res = append(res, []int{nums[item1[0]], nums[item1[1]], nums[item2[0]], nums[item2[1]]})
                        fmt.Println(item1[0], item1[1], item2[0], item2[1], midNum)
                        resMap[h] = true
                    }
                }
            }
        }
    }
    return res
}
