package dushengchen

/**
Submission:
   https://leetcode.com/submissions/detail/740165619/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Right Side View.
Memory Usage: 2.3 MB, less than 77.34% of Go online submissions for Binary Tree Right Side View.
*/
func rightSideView(root *TreeNode) []int {
	return _rightSideView(root, 0, []int{})
}

func _rightSideView(root *TreeNode, deep int, rightView []int) []int {
	if root == nil {
		return rightView
	}
	if deep >= len(rightView) {
		rightView = append(rightView, root.Val)
	} else {
		rightView[deep] = root.Val
	}
	rightView = _rightSideView(root.Left, deep+1, rightView)
	rightView = _rightSideView(root.Right, deep+1, rightView)
	return rightView
}