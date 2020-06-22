package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/356998480/
*/
var digitMaps = map[rune][]rune {
    '2' : []rune{'a', 'b', 'c'},
    '3' : []rune{'d', 'e', 'f'},
    '4' : []rune{'g', 'h', 'i'},
    '5' : []rune{'j', 'k', 'l'},
    '6' : []rune{'m', 'n', 'o'},
    '7' : []rune{'p', 'q', 'r', 's'},
    '8' : []rune{'t', 'u', 'v'},
    '9' : []rune{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
    var res []string
    for _, i := range goDeep([]rune(digits)) {
        res = append(res, string(i))
    }

    return res
}

func goDeep(nextDigits []rune) [][]rune {
    if len(nextDigits) == 0 {
        return nil
    } else if len(nextDigits) == 1 {
        var res [][]rune
        for _, next := range digitMaps[nextDigits[0]] {
            res = append(res, []rune{next})
        }
        return res
    }
    subs := goDeep(nextDigits[1:])
    var res [][]rune
    for _, next := range digitMaps[nextDigits[0]] {
        for _, sub := range subs {
            x := make([]rune, len(sub)+1)
            x[0] = next
            for i := range sub {
                x[i+1] = sub[i]
            }
            res = append(res, x)
        }
    }
    return res
}