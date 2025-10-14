package template

import "athena_go/algorithm/common"

type MaxDepth struct {
}

func (m *MaxDepth) maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	curDepth := 1
	leftDepth := m.maxDepth(root.Left)
	rightDepth := m.maxDepth(root.Right)
	if leftDepth > rightDepth {
		curDepth += leftDepth
	} else {
		curDepth += rightDepth
	}
	return curDepth
}
