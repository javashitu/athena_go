package leetcodetop100

type FindDuplicate struct{}

func (f *FindDuplicate) findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	pre1 := 0
	pre2 := slow
	for pre1 != pre2 {
		pre1 = nums[pre1]
		pre2 = nums[pre2]
	}
	return pre1
}
