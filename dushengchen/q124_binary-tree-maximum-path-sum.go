package dushengchen

/**
Submission:
   https://leetcode.com/submissions/detail/740152233/

Runtime: 40 ms, faster than 11.76% of Go online submissions for Binary Tree Maximum Path Sum.
Memory Usage: 7.9 MB, less than 36.63% of Go online submissions for Binary Tree Maximum Path Sum.
*/
func maxPathSum(root *TreeNode) int {
	pass, max := _maxPathSum(root)
	return MaxInt(pass, max)
}

/**
	返回值为：
	1. 通过该节点的最大值
	2. 该节点子树内的最大值
 */
func _maxPathSum(root *TreeNode) (int, int) {
	if root == nil {
		return -1001, -1001
	}
	lpass, lmax := _maxPathSum(root.Left)
	rpass, rmax := _maxPathSum(root.Right)
	passMax := MaxInt(lpass+root.Val, rpass+root.Val, root.Val)
	return passMax, MaxInt(lmax, rmax, passMax,  lpass+root.Val+rpass)
}
