package dushengchen

/*
Submission:
   https://leetcode.com/submissions/detail/369624404/
*/
func tribonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n <= 2 {
		return 1
	}
	n0 := 0
	n1 := 1
	n2 := 1
	for i := 3; i < n; i++ {
		n2, n1, n0 = n1+n2+n0, n2, n1
	}
	return n2 + n1 + n0
}
