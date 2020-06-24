package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/357699373/
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    endN := make([]*ListNode, n+1)
    curPos := 0
    curPt := head
    for {
        if curPt == nil {
            break
        }
        endN[curPos] = curPt
        curPt = curPt.Next
        curPos++
        if curPos >= len(endN) {
            curPos = 0
        }
    }
    nodeNPre := curPos
    nodeNNext := curPos + 2
    if nodeNNext >= len(endN) {
        nodeNNext = nodeNNext - len(endN)
    }
    if endN[nodeNPre] == nil {
        if nodeNNext < 0 || nodeNNext >= len(endN) {
            return nil
        }
        return endN[nodeNNext]
    } else {
        if nodeNNext < 0 || nodeNNext >= len(endN) || nodeNPre == nodeNNext {
            endN[nodeNPre].Next = nil
        } else {
            endN[nodeNPre].Next = endN[nodeNNext]
        }
    }

    return head
}
