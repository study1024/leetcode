package dushengchen

/*
question:
    https://leetcode.com/problems/container-with-most-water/
Submission:
    https://leetcode.com/submissions/detail/291759879/
*/

func maxArea(a []int) int {
	max, i, j := 0, 0, len(a)-1
	for {
		s := MinInt(a[i], a[j]) * (j - i)
		if s > max {
			max = s
		}
		if a[i] > a[j] {
			j--
		} else {
			i++
		}
		if i == j {
			break
		}
	}
	return max
}
