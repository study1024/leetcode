package dushengchen

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
		if otherPoss, ok := a[other]; ok {
			return returnByRrder(pos, otherPoss)
		}
	}
	return nil
}

func returnByRrder(a int, b []int) []int {
	if a == b[0] {
		return []int{a, b[1]}
	}
	return []int{a, b[0]}
}
