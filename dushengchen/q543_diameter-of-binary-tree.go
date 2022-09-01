package dushengchen

/**
Submission:
   https://leetcode.com/submissions/detail/740781943/

Runtime: 6 ms, faster than 67.71% of Go online submissions for Diameter of Binary Tree.
Memory Usage: 4.7 MB, less than 5.00% of Go online submissions for Diameter of Binary Tree.
*/

func diameterOfBinaryTree(root *TreeNode) int {
	_, b := _diameterOfBinaryTree(root)
	return b
}

func _diameterOfBinaryTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lpass, lmax := _diameterOfBinaryTree(root.Left)
	rpass, rmax := _diameterOfBinaryTree(root.Right)

	return MaxInt(lpass+1, rpass+1), MaxInt(lmax, rmax, lpass+rpass+1)
}