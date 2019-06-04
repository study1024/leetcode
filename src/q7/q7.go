package q7


func reverse(x int) int {
    if x == 0 {
        return 0
    } else if x < 0 {
        return -reverse(-x)
    }

    max := 1<<31 - 1
    res := 0
    for {
        res = res * 10+ x % 10
        x = x / 10

        if res > max {
            return 0
        }
        if x == 0 {
            break
        }
    }

    return res
}

