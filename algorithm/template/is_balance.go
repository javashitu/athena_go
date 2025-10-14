package template

import "athena_go/algorithm/common"

type IsBalance struct {
	balanceFlag bool
}

func (i *IsBalance) isBalance(root *common.TreeNode) bool {
	i.balanceFlag = true
	i.findDepth(root)
	return i.balanceFlag
}

func (i *IsBalance) findDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := i.findDepth(root.Left)
	rightDepth := i.findDepth(root.Right)
	if leftDepth > rightDepth {
		i.balanceFlag = i.balanceFlag && ((leftDepth - rightDepth) <= 1)
		return leftDepth + 1
	} else {
		i.balanceFlag = (i.balanceFlag && (rightDepth-leftDepth) <= 1)
		return rightDepth + 1
	}
}
