package leetcodetop100

type MinDistance struct{}

func (m *MinDistance) minDistance(word1 string, word2 string) int {
	//dp[i][j]表示长度为i的word1转换到长度为j的word2所需要的最小的步数
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	//当i为0时转换到j长度word2需要添加j个字符
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	//当i和j位置的字符一样时dp[i][j]由dp[i-1][j-1]决定
	//当i和j不一样时,dp[i][j] = dp[i-1][j]+1 和dp[i][j - 1] + 1中两者较小的值来决定
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1:i] == word2[j-1:j] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			minDistance := dp[i-1][j]
			if dp[i][j-1] < minDistance {
				minDistance = dp[i][j-1]
			}
			if dp[i-1][j-1] < minDistance {
				minDistance = dp[i-1][j-1]
			}
			dp[i][j] = minDistance + 1

		}
	}
	return dp[len(word1)][len(word2)]
}
