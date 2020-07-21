package dushengchen


/**
Submission:
    https://leetcode.com/submissions/detail/367209719/
*/
func climbStairs(n int) int {
    if n == 1 {
        return 1
    } else if n == 2 {
        return 2
    }
    n1 := 2
    n2 := 1
    for i := 3; i < n; i++ {
        n1, n2 = n1+n2, n1
    }
    return n1+n2
}