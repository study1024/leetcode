package dushengchen

import "math"

/**
Submission:
    https://leetcode.com/submissions/detail/378695017/
*/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	memo := map[int]int{}
	return coinChangeRev(coins, amount, memo)
}

func coinChangeRev(coins []int, amount int, memo map[int]int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if a, ok := memo[amount]; ok {
		return a
	}
	ret := math.MaxInt64
	for _, c := range coins {
		if x := coinChangeRev(coins, amount-c, memo); x >= 0 {
			ret = MinInt(ret, x+1)
		}
	}
	if ret != math.MaxInt64 {
		memo[amount] = ret
	} else {
		memo[amount] = -1
	}

	return memo[amount]
}
