package q1

func twoSum(nums []int, target int) []int {
    a := make(map[int][]int, len(nums))

    for pos, num := range nums {
        a[num] = append(a[num], pos)
    }
    for pos, num := range nums {
        other := target - num
        if other == num && len(a[other]) <= 1 {
            continue
        }
        if other_poss, ok := a[other]; ok {
            return return_by_order(pos, other_poss)
        }
    }
    return nil
}

func return_by_order(a int, b []int) []int {
    if a == b[0] {
        return []int{a, b[1]}
    }
    return []int{a, b[0]}
}