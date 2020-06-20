package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/282703014/
*/

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    } else if len(strs) == 1 {
        return strs[0]
    }
    idx := 0
    for {
        for i:=1; i< len(strs); i++ {
            if idx >= len(strs[0]) || idx >= len(strs[i]) || strs[i][idx] != strs[0][idx] {
                return strs[0][0:idx]
            }
        }
        idx++
    }

}