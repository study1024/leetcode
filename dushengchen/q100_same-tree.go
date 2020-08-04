package dushengchen

/**
Submission:
	https://leetcode.com/submissions/detail/375724887/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
Memory Usage: 2.1 MB, less than 12.50% of Go online submissions for Same Tree.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if q.Val != p.Val {
		return false
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	return isSameTree(p.Right, q.Right)
}
