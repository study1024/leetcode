package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/362016631/
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k <= 1{
        return head
    }
    kMap := make([]*ListNode, k)
    var kHead *ListNode
    kTail := head
    var newHead *ListNode
    for {
        pos := 0
        for ; pos < k; pos++ {
            kMap[pos] = kTail
            kTail = kTail.Next
            if kTail == nil {
                break
            }
        }
        if pos == k || (pos == (k-1) && kTail == nil) {
            for i:= k-1;i>0; i-- {
                kMap[i].Next=kMap[i-1]
            }
            if kHead != nil {
                kHead.Next = kMap[k-1]
            }
            if newHead == nil {
                newHead = kMap[k-1]
            }
            kHead = kMap[0]
            kMap[0] = nil
        }
        if kTail == nil {
            break
        }
    }
    if kHead != nil {
        kHead.Next = kMap[0]
    }
    if newHead != nil {
        return newHead
    } else {
        return head
    }
}