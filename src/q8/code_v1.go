package q8

func myAtoi(str string) int {
    sn := 0
    num := -1
    intMax := 1<< 31 -1
    for _, ch := range str {


        if ch == '-' {
            if num >= 0 || sn != 0{
                break
            }
            sn = -1
        } else if  ch == '+' {
            if num >= 0 || sn != 0{
                break
            }
            sn = 1
        } else if ch >= '0' && ch <= '9' {
            if num <= 0 {
                num = 0
            }
            tmp :=  num * 10 + int(ch - '0')
            if tmp > intMax {
                if sn >= 0 {
                    return intMax
                }
                return -1 << 31
            }
            num = tmp
        } else if ch == ' ' {
            if num >= 0 || sn != 0{
                break
            }
        } else {
            break
        }
    }
    if sn == 0 {
        sn = 1
    }
    if num <= 0 {
        num = 0
    }
    return num * sn
}
