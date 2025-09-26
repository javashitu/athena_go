package leetcodehot100

import "athena_go/algorithm/common"

type SumTwoNum struct {
}

func (s *SumTwoNum) AddTwoSum(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	dummyHead := &common.ListNode{}
	curNode := dummyHead
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + carry
		carry = sum / 10
		sum = sum % 10
		curNode.Next = &common.ListNode{Val: sum}
		curNode = curNode.Next
	}
	return dummyHead.Next
}
