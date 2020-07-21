package dushengchen


/**
Submission:
    https://leetcode.com/submissions/detail/367236667/
*/
func minCostClimbingStairs(cost []int) int {
    if len(cost) <= 1 {
        return 0
    } else if len(cost) == 2 {
        return MinInt(cost...)
    }
    cost = append(cost, 0)
    costN2 := 0
    costN1 := 0
    costN := 0
    for i:= 2; i < len(cost); i++ {
        costN = MinInt(costN2+cost[i-2], costN1+cost[i-1])
        costN2 = costN1
        costN1 = costN
        // fmt.Println("step=", i, "cost=", costN)
    }
    return costN
}
