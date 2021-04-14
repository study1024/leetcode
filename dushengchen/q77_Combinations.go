package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/480451077/
*/
func combine(n int, k int) [][]int {
	var result [][]int
	if k == 1 {
		for i:=1; i<=n; i++ {
			result = append(result, []int{i})
		}
		return result
	}
	subResult := combine(n, k-1)
	for idx := range subResult {
		for i := subResult[idx][len(subResult[idx])-1]+1; i <= n; i++ {
			x := make([]int, len(subResult[idx])+1)
			copy(x, subResult[idx])
			x[len(subResult[idx])] = i
			result = append(result, x)
		}
	}
	return result
}