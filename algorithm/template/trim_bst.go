package template

import "athena_go/algorithm/common"

type TrimBST struct {
}

func (t *TrimBST) trimBst(root *common.TreeNode, low int, high int) *common.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return t.trimBst(root.Right, low, high)
	}
	if root.Val > high {
		return t.trimBst(root.Left, low, high)
	}
	root.Left = t.trimBst(root.Left, low, high)
	root.Right = t.trimBst(root.Right, low, high)

	return root
}
