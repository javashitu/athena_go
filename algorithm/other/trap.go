package other

type Trap struct{}

func (t *Trap) trap(nums []int) int {
	preArr := make([]int, len(nums))
	postArr := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		preIndex := preArr[i-1]
		if nums[preIndex] > nums[i] {
			preArr[i] = preIndex
		} else {
			preArr[i] = i
		}
	}
	postArr[len(nums)-1] = len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		postIndex := postArr[i+1]
		if nums[postIndex] > nums[i] {
			postArr[i] = postIndex
		} else {
			postArr[i] = i
		}
	}
	result := 0
	for i := 1; i < len(nums)-1; i++ {
		minIndex := 0
		if nums[preArr[i]] > nums[postArr[i]] {
			minIndex = postArr[i]
		} else {
			minIndex = preArr[i]
		}
		if nums[minIndex] > nums[i] {
			result += (nums[minIndex] - nums[i])
		}
	}
	return result
}
