package dushengchen

/**
Submission:
	https://leetcode.com/submissions/detail/375719502/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	x := []int{}
	inorderTraversalSeq(root, &x)
	return x
}

func inorderTraversalSeq(root *TreeNode, seq *[]int) {
	if root == nil {
		return
	}
	inorderTraversalSeq(root.Left, seq)
	*seq = append(*seq, root.Val)
	inorderTraversalSeq(root.Right, seq)
}
