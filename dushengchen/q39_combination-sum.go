package dushengchen


/**
Submission:
    https://leetcode.com/submissions/detail/364143609/
*/

func combinationSum(candidates []int, target int) [][]int {
    candidates = SortInt(candidates)
    return combinationSumRecursive(candidates, target)
}

func combinationSumRecursive(candidates []int, target int) [][]int {
    //fmt.Println("in", candidates, target)
    all := [][]int{}
    for i, c := range candidates {
        rest := target - c
        if rest == 0 {
            all= append(all, []int{c})
            continue
        }
        if rest < candidates[0] {
            continue
        }
        ret := combinationSumRecursive(candidates[0:i+1], rest)
        if len(ret) > 0 {
            for i := range ret {
                ret[i] = append(ret[i], c)
            }
            all = append(all, ret...)
        }
    }
    //fmt.Println("out", candidates, target, all)
    return all
}

//func combinationSum(candidates []int, target int) [][]int {
//    used := map[int]bool{}
//    targetMap := map[int][][]int{}
//    candidates = SortInt(candidates)
//    combinationSumRecursive(candidates, target, &used, &targetMap)
//    // return targetMap[target]
//    x := targetMap[target]
//    resMap := map[string]bool{}
//    res := [][]int{}
//    for i := range x {
//        b := SortInt(x[i])
//        s := fmt.Sprintf("%v", b)
//        if resMap[s] {
//            continue
//        }
//        resMap[s] = true
//        res = append(res, x[i])
//    }
//    return res
//}
//
//func combinationSumRecursive(candidates []int, target int, used *map[int]bool, targetMap *map[int][][]int) {
//    fmt.Println(candidates, target, used, targetMap)
//    targetRes := [][]int{}
//    if _, ok := (*targetMap)[target]; ok {
//        return
//    }
//    if len(candidates) == 0 {
//        return
//    }
//    if target < candidates[0] {
//        return
//    }
//    for i := len(candidates)-1; i >= 0; i-- {
//        c := candidates[i]
//        if (*used)[c] {
//            continue
//        }
//        rest := target - c
//        if rest == 0 {
//            targetRes = append(targetRes, []int{c})
//            continue
//        }
//        if rest < candidates[0] {
//            continue
//        }
//        // (*used)[c] = true
//        combinationSumRecursive(candidates[:i+1], rest, used, targetMap)
//        // (*used)[c] = false
//        if _, ok := (*targetMap)[rest]; ok {
//            for i := range (*targetMap)[rest] {
//                x := []int{}
//                x = append(x, (*targetMap)[rest][i]...)
//                x = append(x, c)
//                targetRes = append(targetRes, x)
//            }
//        }
//    }
//    if len(targetRes) > 0 {
//        (*targetMap)[target] = targetRes
//    }
//    // (*used)[candidates[0]] = true
//}
//
//
//func SortInt(a []int) []int {
//    for i:=0; i<len(a); i++ {
//        for j:=i+1; j<len(a); j++ {
//            if a[i] > a[j] {
//                a[i],a[j] = a[j],a[i]
//            }
//        }
//    }
//    return a
//}