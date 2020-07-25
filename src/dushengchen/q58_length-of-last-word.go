package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/371271670/
*/
func lengthOfLastWord(s string) int {
	// wordMap := map[string]bool{}
	lastWordStart, lastWordEnd := 0, 0
	start := 0
	r := []rune(s)
	for i, v := range r {
		if v == ' ' {
			if start < i {
				lastWordStart, lastWordEnd = start, i
			}
			start = i + 1
		}
	}
	if start <= len(r)-1 {
		lastWordStart, lastWordEnd = start, len(r)
	}
	// fmt.Println(string(r[lastWordStart:lastWordEnd]))
	return lastWordEnd - lastWordStart
}
