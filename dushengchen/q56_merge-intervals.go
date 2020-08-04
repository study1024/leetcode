package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/371219433/
*/

func merge(intervals [][]int) [][]int {
	sortByStart(intervals)
	var ret [][]int
	var cur []int
	for i, v := range intervals {
		if len(cur) == 0 {
			cur = intervals[i]
			continue
		}
		if cur[1] >= v[0] {
			if cur[1] < v[1] {
				cur[1] = v[1]
			}
		} else {
			ret = append(ret, cur)
			cur = intervals[i]
		}
	}
	if len(cur) > 0 {
		ret = append(ret, cur)
	}
	return ret
}

func sortByStart(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i][0] > a[j][0] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
