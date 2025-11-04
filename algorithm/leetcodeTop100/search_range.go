package leetcodetop100

type SearchRange struct{}

func (s SearchRange) searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := s.getLeft(nums, target)
	right := s.getRight(nums, target)
	return []int{left, right}
}

func (s SearchRange) getLeft(nums []int, target int) int {
	low := 0
	high := len(nums)
	for low < high {
		middle := low + (high-low)/2
		if nums[middle] < target {
			low = middle + 1
		} else if nums[middle] > target {
			high = middle
		} else {
			high = middle
		}
	}
	if high <= len(nums)-1 && nums[high] == target {
		return high
	}
	return -1
}
func (s SearchRange) getRight(nums []int, target int) int {
	low := 0
	high := len(nums)
	for low < high {
		middle := low + (high-low)/2
		if nums[middle] < target {
			low = middle + 1
		} else if nums[middle] > target {
			high = middle
		} else {
			low = middle + 1
		}
	}
	if low != 0 && nums[low-1] == target {
		return low - 1
	}
	return -1
}
