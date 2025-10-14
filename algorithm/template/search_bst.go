package template

import "athena_go/algorithm/common"

type SearchBST struct {
}

func (s *SearchBST) searchBst(root *common.TreeNode, val int) *common.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return s.searchBst(root.Left, val)
	} else {
		return s.searchBst(root.Right, val)
	}
}
