package template

import "sort"

type FindContentChild struct{}

func (f *FindContentChild) findContentChild(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	j := len(g) - 1
	for i := len(s) - 1; i >= 0 && j >= 0; {
		if s[i] >= g[j] {
			count++
			i--
		}
		j--

	}
	return count
}
