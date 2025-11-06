package leetcodetop100

import "sort"

type GroupAnagrams struct{}

func (g *GroupAnagrams) groupAnagrams(arr []string) [][]string {
	hash := make(map[string][]string)
	for i := 0; i < len(arr); i++ {
		str := arr[i]
		tempArr := []byte(str)
		sort.Slice(tempArr, func(i, j int) bool {
			return tempArr[i] < tempArr[j]
		})
		temp := string(tempArr)
		hash[temp] = append(hash[temp], arr[i])
	}
	result := make([][]string, 0)
	for _, value := range hash {
		result = append(result, value)
	}
	return result
}
