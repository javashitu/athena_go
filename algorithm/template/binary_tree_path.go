package template

import (
	"athena_go/algorithm/common"
	"strconv"
)

type BinaryTreePath struct {
}

func (b *BinaryTreePath) binaryTreePath(root *common.TreeNode) []string {
	result := []string{}
	path := []int{}
	b.findPath(root, &path, &result)
	return result
}

func (b *BinaryTreePath) findPath(root *common.TreeNode, path *[]int, result *[]string) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	if root != nil && root.Left == nil && root.Right == nil {
		tempPath := ""
		for index, val := range *path {
			if index == len(*path)-1 {
				tempPath += strconv.Itoa(val)
			} else {
				tempPath += strconv.Itoa(val) + "->"
			}
		}
		*result = append(*result, tempPath)
	}
	if root.Left != nil {
		b.findPath(root.Left, path, result)
		*path = (*path)[:len(*path)-1]
	}
	if root.Right != nil {
		b.findPath(root.Right, path, result)
		*path = (*path)[:len(*path)-1]
	}
}
