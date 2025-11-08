package leetcodetop100

import "athena_go/algorithm/common"

type Flatten struct {
	// DummyHead *common.TreeNode
	// PreNode   *common.TreeNode
}

func (f *Flatten) flatten(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	leftNode := f.flatten(root.Left)
	rightNode := f.flatten(root.Right)
	root.Left = nil
	if leftNode != nil {
		root.Right = leftNode
		for leftNode.Right != nil {
			leftNode = leftNode.Right
		}
		leftNode.Right = rightNode
	} else {
		root.Right = rightNode
	}
	return root
}
