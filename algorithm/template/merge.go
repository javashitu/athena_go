package template

import "sort"

type Merge struct{}

func (m *Merge) merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left := intervals[0][0]
	right := intervals[0][1]
	result := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		curLeft := intervals[i][0]
		curRight := intervals[i][1]
		if curLeft > right {
			result = append(result, []int{left, right})
			left = curLeft
			right = curRight
		} else {
			if curRight > right {
				right = curRight
			}
		}
	}
	result = append(result, []int{left, right})
	return result
}
