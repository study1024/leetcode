package main

import "leetcode/dushengchen"

func main() {
	//dushengchen.LargestRectangleArea([]int{12,11,10,9,8,7,6,5,4,3,2,1})、
	head := &dushengchen.ListNode{}
	cur := head
	for _, v := range []int{1,4,3,2,5,2} {
		cur.Next = &dushengchen.ListNode{Val: v}
		cur = cur.Next
	}
	dushengchen.Partition(head.Next, 3)
}
