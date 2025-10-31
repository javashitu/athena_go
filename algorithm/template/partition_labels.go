package template

type PartitionLabels struct{}

func (p *PartitionLabels) partitionLabels(str string) []int {
	hash := make([]int, 26)
	for i := 0; i < len(str); i++ {
		index := str[i] - 'a'
		hash[index] = i
	}
	result := []int{}
	left := 0
	right := 0
	for i := 0; i < len(str); i++ {
		index := hash[str[i]-'a']
		if index > right {
			right = index
		}
		if right == i {
			result = append(result, right-left+1)
			left = i + 1
		}

	}
	return result
}
