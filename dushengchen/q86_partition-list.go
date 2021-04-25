package dushengchen

/*
Submission:
	https://leetcode.com/submissions/detail/484823016/
*/
func Partition(head *ListNode, x int) *ListNode {
	var mHead *ListNode
	p := &ListNode{}
	subHead := &ListNode{}
	cur := subHead
	p.Next = head
	for {
		if p == nil || p.Next == nil {
			break
		}
		if p.Next.Val < x {
			cur.Next = p.Next
			cur = cur.Next
			p.Next = p.Next.Next
			continue
		} else if mHead == nil {
			mHead = p.Next
		}
		p = p.Next
	}
	cur.Next = mHead
	return subHead.Next
}