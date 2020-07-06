package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/362016631/
*/

func swapPairs(head *ListNode) *ListNode {
    return reverseKGroup(head, 2)
}