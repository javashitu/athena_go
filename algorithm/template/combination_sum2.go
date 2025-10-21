package template

import "sort"

type CombinationSum2 struct {
	result [][]int
}

func (c *CombinationSum2) CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	c.backtracking(candidates, target, 0, 0, make([]bool, len(candidates)), make([]int, 0))
	return c.result
}

func (c *CombinationSum2) backtracking(candidates []int, target int, startIndex int, sum int, used []bool, arr []int) {
	if sum == target {
		temp := append(make([]int, 0), arr...)
		c.result = append(c.result, temp)
		return
	}
	if sum > target {
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		sum += candidates[i]
		arr = append(arr, candidates[i])
		used[i] = true
		if (i > 0 && candidates[i] == candidates[i-1] && used[i-1]) || (i == 0) || (candidates[i] != candidates[i-1]) {
			c.backtracking(candidates, target, i+1, sum, used, arr)
		}
		sum -= candidates[i]
		arr = arr[0 : len(arr)-1]
		used[i] = false
	}
}
