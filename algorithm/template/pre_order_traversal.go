package template

import common "athena_go/algorithm/common"

type PreOrder struct {
}

func (p *PreOrder) preorderTraversal(root *common.TreeNode) []int {
	arr := []int{}
	var inOrder func(node *common.TreeNode)
	inOrder = func(node *common.TreeNode) {
		if node == nil {
			return
		}
		arr = append(arr, node.Val)
		inOrder(node.Left)
		inOrder(node.Right)
	}
	inOrder(root)
	return arr
}

func (p *PreOrder) preorder2(root *common.TreeNode) []int {
	vals := []int{}
	stack := []*common.TreeNode{}
	if root == nil {
		return vals
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		vals = append(vals, node.Val)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return vals
}
