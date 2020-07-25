package dushengchen


/*
Submission:
    https://leetcode.com/submissions/detail/356117314/
*/

type chatIntPar struct {
    Char string
    Val int
}

var allRomanChar = []chatIntPar{
    chatIntPar{"M", 1000},
    chatIntPar{"CM", 900},
    chatIntPar{"D", 500},
    chatIntPar{"CD", 400},

    chatIntPar{"C", 100},
    chatIntPar{"XC", 90},
    chatIntPar{"L", 50},
    chatIntPar{"XL", 40},

    chatIntPar{"X", 10},
    chatIntPar{"IX", 9},
    chatIntPar{"V", 5},
    chatIntPar{"IV", 4},

    chatIntPar{"III", 3},
    chatIntPar{"II", 2},
    chatIntPar{"I", 1},
}


func intToRoman(num int) string {
    res := ""
    for {
        if num <= 0 {
            break
        }
        for _, c := range allRomanChar {
            if num >= c.Val {
                num, res = num - c.Val, res+c.Char
                break
            }
        }
    }
    return res
}