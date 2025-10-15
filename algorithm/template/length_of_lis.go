package template

type LengthOfLis struct{}

func (l *LengthOfLis) lengthOfLis(arr []int) int {
	// 1. dp[i]表示严格以i结尾的递增数组长度
	//2. dp[i]
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}
	maxLen := 1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = l.max(dp[i], dp[j]+1)
			}
			maxLen = l.max(maxLen, dp[i])
		}
	}
	return maxLen
}

func (l *LengthOfLis) max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
