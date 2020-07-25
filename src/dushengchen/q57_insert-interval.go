package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/371249931/
*/

// https://leetcode.com/submissions/detail/371249931/
//O(n^2)的复杂度
// func insert(intervals [][]int, newInterval []int) [][]int {
// 	intervals = append(intervals, newInterval)
// 	return merge(intervals)
// }

//O(n)的复杂度
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var ret [][]int
	var cur []int
	find := false
	ret = append(ret, newInterval)
	for i := 0; i < len(intervals); i++ {
		// fmt.Println(i, cur, ret)
		if len(cur) == 0 {
			if intervals[i][0] < newInterval[0] && intervals[i][1] < newInterval[0] {
				ret = append(ret, intervals[i])
				continue
			}
			find = true
			if intervals[i][0] > newInterval[0] {
				cur = newInterval
			} else {
				cur = intervals[i]
				intervals[i] = newInterval
			}
		}
		if cur[1] >= intervals[i][0] {
			if cur[1] < intervals[i][1] {
				cur[1] = intervals[i][1]
			}
		} else {
			ret = append(ret, cur)
			cur = intervals[i]
		}
	}
	if len(cur) > 0 {
		ret = append(ret, cur)
	}
	// fmt.Println(find, ret)
	if find {
		return ret[1:]
	}
	if newInterval[0] < ret[1][0] {
		return ret
	}
	ret = append(ret, newInterval)
	return ret[1:]
}
