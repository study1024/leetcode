package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/282703014/
*/

import "fmt"

func threeSum(nums []int) [][]int {
    nums = SortInt(nums)
    var find int
    var res [][]int
    resMap := map[string]bool{}
    numsMap := map[int]int{}
    for k,v :=range nums {
        numsMap[v] = k
    }
    for i:=0; i<len(nums); i++ {
        for j:=i+1; j<len(nums); j++ {
            find = 0-(nums[i]+nums[j])
            if pos, ok := numsMap[find]; !ok || pos <= j {
                continue
            }
            key := fmt.Sprintf("%d:%d:%d", nums[i],find,nums[j])
            if !resMap[key] {
                resMap[key] = true
                res = append(res, []int{nums[i],find,nums[j]})
            }
        }
    }
    return res
}

//
//func sort(a []int) []int {
//    for i:=0; i<len(a); i++ {
//        for j:=i+1; j<len(a); j++ {
//            if a[i] > a[j] {
//                a[i],a[j] = a[j],a[i]
//            }
//        }
//    }
//    return a
//}