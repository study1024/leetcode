package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/282698341/
*/

var roman2intMap = map[byte]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}
func romanToInt(s string) int {

    ret := 0
    for i:=0; i < len(s); i++ {
        num, ok := roman2intMap[s[i]]
        if !ok {
            break
        }
        if (i+1) < len(s) {
            numNext, ok := roman2intMap[s[i+1]]
            if !ok {
                break
            }
            if numNext > num {
                ret = ret - num
                continue
            }
        }
        ret = ret + num
    }
    return ret
}
