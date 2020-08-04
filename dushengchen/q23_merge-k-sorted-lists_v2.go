package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/357219769/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKListsV2(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    cur := lists[0]
    for i:=1; i < len(lists); i++ {
        cur = mergeTwoLists(cur, lists[i])
    }
    return cur
}