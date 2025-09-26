package leetcodehot100

import "athena_go/algorithm/common"

type ListHasCycle struct {
}

func (l *ListHasCycle) HasCycle(head *common.ListNode) bool {
	fast := head
	slow := head

	for slow != nil && slow.Next != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
