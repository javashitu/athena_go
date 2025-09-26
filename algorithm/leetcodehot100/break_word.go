package leetcodehot100

type WordBreak struct {
}

func (w *WordBreak) BreakWord(s string, wordDict []string) bool {
	//1. dp数组的含义,dp[i]表示前i个[0,i)字符能不能全部被wordDict离得单词组成
	dp := make([]bool, len(s)+1)
	dp[0] = true
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
