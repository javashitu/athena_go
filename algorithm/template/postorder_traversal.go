package template

import common "athena_go/algorithm/common"

type PostOrderTraversal struct {
}

func (p *PostOrderTraversal) PostOrder(root *common.TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*common.TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		result = append([]int{node.Val}, result...)
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return result
}
