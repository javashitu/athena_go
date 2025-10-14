package template

import "athena_go/algorithm/common"

type LowestCommonAncestor struct {
}

func (l *LowestCommonAncestor) lowestCommonAncestor(root *common.TreeNode, p, q *common.TreeNode) *common.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := l.lowestCommonAncestor(root.Left, p, q)
	right := l.lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
}
