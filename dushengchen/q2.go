package dushengchen

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    cur := head
    nextDigit := 0
    for {
        cur.Val = GetVal(l1) + GetVal(l2) + nextDigit
        nextDigit = cur.Val / 10
        cur.Val = cur.Val % 10

        l1 = next(l1)
        l2 = next(l2)

        if l1 == nil && l2 == nil && nextDigit == 0 {
            break
        }
        cur.Next = &ListNode{}
        cur = cur.Next
    }
    return head
}

func next(l *ListNode) *ListNode {
    if l == nil {
        return nil
    }
    return l.Next
}

func GetVal(l *ListNode) int {
    if l == nil {
        return 0
    }
    return l.Val
}
