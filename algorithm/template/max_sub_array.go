package template

import "math"

type MaxSubArray struct{}

func (m *MaxSubArray) maxSubArray(nums []int) int {
	sum := 0
	// left := 0
	maxSum := math.MinInt
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
			continue
		}
	}

	return maxSum
}
