package template

import "athena_go/algorithm/common"

type MinCameraConv struct {
	count int
}

func (m *MinCameraConv) minCameraConv(root *common.TreeNode) int {
	if m.traversal(root) == 0 {
		m.count++
	}
	return m.count
}

//0没有被覆盖，1使用了监视器，2被覆盖到了
func (m *MinCameraConv) traversal(root *common.TreeNode) int {
	if root == nil {
		return 2
	}
	left := m.traversal(root.Left)
	right := m.traversal(root.Right)
	//这一行代码可以说是很巧妙的精华处理，必须先考虑叶子结点的状态
	if left == 2 && right == 2 {
		return 0
	}
	//左右子树任意一个没有被覆盖，那么就需要使用监视器
	if left == 0 || right == 0 {
		m.count++
		return 1
	}
	//左右子树都被覆盖了，并且有一个子树使用了监视器，当前节点能够被覆盖，不使用监视器
	if left == 1 || right == 1 {
		return 2
	}
	return -1
}
