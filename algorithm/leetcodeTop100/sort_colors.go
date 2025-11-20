package leetcodetop100

type SortColors struct{}

func (s *SortColors) sortColors(nums []int) []int {
	index0 := 0
	index1 := 0
	index2 := 0
	//分为3个片段，每个片段由指针指向这个片段的结束位置
	for i := range nums {
		if nums[i] == 0 {
			nums[index2] = 2
			index2++
			nums[index1] = 1
			index1++
			nums[index0] = 0
			index0++
		} else if nums[i] == 1 {
			nums[index2] = 2
			index2++
			nums[index1] = 1
			index1++
		} else {
			nums[index2] = 2
			index2++
		}
	}
	return nums
}
