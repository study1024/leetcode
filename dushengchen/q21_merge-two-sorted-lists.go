package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/270874219/
*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head :=  &ListNode{}
    cur := head
    for {
        if l1 == nil {
            cur.Next = l2
            break
        }
        if l2 == nil {
            cur.Next = l1
            break
        }
        if l1.Val < l2.Val {
            cur.Next = l1
            l1 = l1.Next
            cur = cur.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
            cur = cur.Next
        }
    }
    return head.Next
}