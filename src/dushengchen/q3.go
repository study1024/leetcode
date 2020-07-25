package dushengchen


func lengthOfLongestSubstring(a string) int {
    if len(a) == 0 {
        return 0
    }
    curStart, curEnd, max := 0, 0, 0
    charMap := map[rune]int{}
    for idx, ch := range a {
        if idxBefore, ok := charMap[ch]; ok && idxBefore >= curStart  {
            curStart = idxBefore+1
            curEnd = idx
        }  else {
            curEnd = idx
            if curEnd - curStart > max {
                max = curEnd - curStart
            }
        }
        charMap[ch] = idx
    }
    return max+1
}
