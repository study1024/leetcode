package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/359411013/
*/

func strStr(haystack string, needle string) int {
    if haystack == "" && needle == "" {
        return 0
    }
    h := []rune(haystack)
    for i := range h {
        if i + len(needle) > len(h) {
            break
        }
        find := true
        for j, y := range needle {
            if (i+j) >= len(h) || h[i+j] != y {
                find = false
                break
            }
        }
        if find {
            return i
        }
    }
    return -1
}