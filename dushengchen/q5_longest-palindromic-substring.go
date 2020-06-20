package dushengchen

/*
question:
    https://leetcode.com/problems/longest-palindromic-substring/
Submission:
    https://leetcode.com/submissions/detail/289234911/

*/

func longestPalindrome(s string) string {
    var max string
    for i:=0; i < len(s); i++ {
        start, end := i, i
        for ;(end+1 < len(s)) && s[start] == s[end+1]; end++ {

        }
        i = end
        x := 1
        for {
            if (start - x) >= 0 && (end + x) < len(s) && s[start-x] == s[end+x] {
                x++
            }else {
                break
            }
        }
        if (end - start + 1 + 2 * (x-1)) > len(max) {

            max = s[start-x+1: end+x]
            //fmt.Printf("start=%d, end=%d, max=%s, x=%d", start, end, max, x)
        }
    }
    return max
}