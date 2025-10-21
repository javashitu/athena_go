package template

import "fmt"

type CombinationSum struct {
	result [][]int
}

func (c *CombinationSum) CombinationSum2(candidates []int, target int) [][]int {
	c.result = [][]int{}
	c.backtrack(candidates, target, 0, 0, make([]int, 0))
	fmt.Printf("finally result %v \n", c.result)
	return c.result
}

func (c *CombinationSum) backtrack(candidates []int, target int, startIndex int, sum int, arr []int) {
	fmt.Printf("the sum is %d the arr is %v \n", sum, arr)
	if sum == target {
		temp := make([]int, 0)
		temp = append(temp, arr...)
		c.result = append(c.result, temp)
		fmt.Printf("the result is %v  \n", c.result)
		return
	}
	if sum > target {
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		sum += candidates[i]
		arr = append(arr, candidates[i])
		c.backtrack(candidates, target, i, sum, arr)
		sum -= candidates[i]
		arr = arr[:len(arr)-1]
	}
}
