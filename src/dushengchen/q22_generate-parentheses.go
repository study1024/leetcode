package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/270874219/
*/

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	res := make([]map[string]bool, n+1)
	res[0] = map[string]bool{}
	res[1] = map[string]bool{"()": true}
	for i := 2; i <= n; i++ {
		resMap := map[string]bool{}
		for v := range res[i-1] {
			resMap[v+"()"] = true
			resMap["()"+v] = true
			resMap["("+v+")"] = true
		}
		for j := 2; j <= i-j; j++ {
			for x := range res[j] {
				for y := range res[i-j] {
					resMap[x+y] = true
					resMap[y+x] = true
				}
			}
		}
		res[i] = resMap
	}
	ret := []string{}
	for k := range res[n] {
		ret = append(ret, k)
	}
	return ret
}
