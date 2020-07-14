package dushengchen


/**
Submission:
    https://leetcode.com/submissions/detail/366393972/
*/
func multiply(num1 string, num2 string) string {
    a := []rune(num1)
    b := []rune(num2)
    sum := make([]rune, len(num1)+len(num2))
    sum[0] = '0'
    for i := 0; i < len(a); i++ {
        m := runeToInt(a[len(a)-i-1])
        for j := 0; j < len(b); j++ {
            n := runeToInt(b[len(b)-j-1])
            x :=  m * n
            for add := 0;; add++{
                c := runeToInt(sum[i+j+add])
                if c > 9 || c < 0 {
                    c = 0
                }
                x = x+c
                rest := x % 10
                sum[i+j+add] = intToRune(rest)
                x = x / 10
                // fmt.Println(m, n, c, x, string(sum))
                if  x == 0 {
                    break
                }
            }
        }
        // fmt.Println(i, m, string(sum))
    }
    ret := []rune{}
    allZero := true
    for i := cap(sum)-1; i >=0; i-- {
        if m := runeToInt(sum[i]); m >= 0 && m <= 9 {
            ret = append(ret, sum[i])
            if m != 0 {
                allZero = false
            }
        }
    }
    if allZero {
        return "0"
    }
    return string(ret)
}

func runeToInt(c rune) int {
    return int(c - '0')
}

func intToRune(i int) rune {
    return rune(i) + '0'
}