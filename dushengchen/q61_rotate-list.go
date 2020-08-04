package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/371282799/
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	kStart, kEnd := head, head
	for i := 0; i < k; i++ {
		if kEnd.Next != nil {
			kEnd = kEnd.Next
		} else {
			kEnd = head
		}
	}
	if kStart == kEnd {
		return head
	}
	// fmt.Println(kStart.Val, kEnd.Val)
	for kEnd.Next != nil {
		kStart, kEnd = kStart.Next, kEnd.Next
	}
	// fmt.Println(kStart.Val, kEnd.Val, head.Val)
	head, kStart.Next, kEnd.Next = kStart.Next, nil, head
	return head
}
