package dushengchen

import "fmt"

//MinInt MinInt
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

//MaxInt MaxInt
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

//MinFloat64 MinFloat64
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

//MaxFloat64 MaxFloat64
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

//ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

//SortInt SortInt
func SortInt(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

//AbsInt AbsInt
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//HashInt HashInt
func HashInt(arg ...int) string {
	n := SortInt(arg)
	return fmt.Sprintf("%#v", n)
}

//NewIntArray2D 新建一个二维数组
func NewIntArray2D(m int, n int) [][]int {
	x := [][]int{}
	for i := 0; i < m; i++ {
		x = append(x, make([]int, n))
	}
	return x
}

//RevertInt 翻转一个数组
func RevertInt(a []int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}
