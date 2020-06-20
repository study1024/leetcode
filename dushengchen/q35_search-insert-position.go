package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/270874219/
*/


func searchInsert(nums []int, target int) int {
    insert := 0
    for idx, v := range nums {
        if v >= target {
            break
        } else {
            insert = idx+1
        }
    }
    return insert;
}