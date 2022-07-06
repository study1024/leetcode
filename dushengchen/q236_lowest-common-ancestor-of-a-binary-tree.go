package dushengchen
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
https://leetcode.com/submissions/detail/740046542/

Runtime: 20 ms, faster than 32.24% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
Memory Usage: 7.1 MB, less than 81.52% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if pf, qf, n := isFather(root, p, q); pf && qf {
		return n
	}
	return root
}

func isFather(root, p, q *TreeNode) (bool, bool, *TreeNode) {
	if root == nil || q == nil || p == nil {
		return false, false, nil
	}
	pf1, qf1, n := isFather(root.Left, p, q)
	if pf1 && qf1 {
		return pf1 ,qf1, n
	}
	pf2, qf2, n := isFather(root.Right, p, q)
	if pf2 && qf2 {
		return pf2 ,qf2, n
	}
	pf := pf1 || pf2 || root.Val == p.Val
	qf := qf1 || qf2 || root.Val == q.Val
	if pf && qf {
		return true, true, root
	}
	return pf, qf, nil
}