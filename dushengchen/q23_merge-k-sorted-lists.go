package dushengchen

import "fmt"

/*
Submission:
    https://leetcode.com/submissions/detail/357218625/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    head :=  &ListNode{}
    cur := head
    for {
        min := -1
        for i := range lists {
            if lists[i] == nil {
                continue
            }
            if min == -1 || lists[i].Val < lists[min].Val {
                min = i
            }
        }
        if min == -1 {
            break
        }
        fmt.Println(min, lists[min].Val)
        cur.Next = lists[min]
        cur = cur.Next
        lists[min] = lists[min].Next
    }
    return head.Next
}