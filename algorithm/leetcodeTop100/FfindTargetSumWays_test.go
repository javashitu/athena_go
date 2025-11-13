package leetcodetop100

type FindTargetSumWays struct{}

func (f *FindTargetSumWays) findTargetSumWays(nums []int, target int) int {
	//2sum(p) = target + sum(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if (target+sum)%2 == 1 || target > sum || f.abs(sum) < f.abs(target) {
		return 0
	}
	sum = (target + sum) / 2
	if sum < 0 {
		return 0
	}
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[sum]
}

func (f *FindTargetSumWays) abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if (sum+target)%2 == 1 || sum < target {
		return 0
	}
	result := (sum + target) / 2
	dp := make([]int, result+1)
	dp[0] = 1
	for num := range nums {
		for i := result; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[result]
}
