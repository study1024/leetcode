package dushengchen


func MinInt(args ...int) int {
    if len(args) == 0 {
        return -1
    }
    min := args[0]
    for _, x := range args {
        if min > x {
            min = x
        }
    }
    return min
}

func MaxInt(args ...int) int {
    if len(args) == 0 {
        return -1
    }
    max := args[0]
    for _, x := range args {
        if max < x {
            max = x
        }
    }
    return max
}

func MinFloat64(args ...float64) float64 {
    if len(args) == 0 {
        return -1
    }
    min := args[0]
    for _, x := range args {
        if min > x {
            min = x
        }
    }
    return min
}


func MaxFloat64(args ...float64) float64 {
    if len(args) == 0 {
        return -1
    }
    max := args[0]
    for _, x := range args {
        if max < x {
            max = x
        }
    }
    return max
}

type ListNode struct {
    Val int
    Next *ListNode
}

func SortInt(a []int) []int {
    for i:=0; i<len(a); i++ {
        for j:=i+1; j<len(a); j++ {
            if a[i] > a[j] {
                a[i],a[j] = a[j],a[i]
            }
        }
    }
    return a
}

func AbsInt(a int) int {
   if a < 0 {
       return -a
   }
   return a
}