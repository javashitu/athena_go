package template

import common "athena_go/algorithm/common"

type InOrder struct {
}

func (i *InOrder) InorderTraversal(root *common.TreeNode) []int {
	vals := []int{}
	stack := []*common.TreeNode{}
	if root == nil {
		return vals
	}
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			vals = append(vals, node.Val)
			root = node.Right
		}
	}
	return vals
}
