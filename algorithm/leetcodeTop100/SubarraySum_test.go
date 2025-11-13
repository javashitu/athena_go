package leetcodetop100

type SubArraySum struct{}

func (s *SubArraySum) subarraySum(nums []int, target int) int {
	prefixMap := make(map[int]int)
	prefixMap[0] = 1
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		index := prefixSum[i] - target
		if value, exists := prefixMap[index]; exists {
			count += value
		}
		prefixMap[prefixSum[i]]++
	}
	return count
}

func (s *SubArraySum) subArraySum2(nums []int, target int) int {
	preArr := make([]int, len(nums))
	for index := range nums {
		if index == 0 {
			preArr[index] = nums[index]
		} else {
			preArr[index] = nums[index] + preArr[index-1]
		}
	}
	count := 0
	remember := make(map[int]int)
	remember[0] = 1
	for index := range preArr {
		remember[preArr[index]]++
		diff := preArr[index] - target
		count += remember[diff]
	}
	return count
}
