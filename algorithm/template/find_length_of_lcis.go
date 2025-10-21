package template

type FindLengthOfLcis struct{}

func (f *FindLengthOfLcis) findLengthOfLcis(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if dp[i] > maxLength {
				maxLength = dp[i]
			}
		}
	}
	return maxLength
}
