package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/357703992/
*/

func isValid(s string) bool {
    stack := make([]rune, len(s))
    pos := 0
    for _, ch := range s {
        switch ch {
        case '(', '{', '[':
            stack[pos] = ch
            pos++
        case ')', '}', ']':
            if pos <= 0 || opponent(stack[pos -1]) != ch{
                return false
            }
            pos--
        }
    }
    if pos == 0 {
        return true
    }
    return false
}

func opponent(ch rune) rune {
    switch ch{
    case '(':
        return ')'
    case '{':
        return '}'
    case '[':
        return ']'
    }
    return' '
}

