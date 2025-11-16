package leetcodetop100

type RemoveInvalidParentheses struct{}

func (r *RemoveInvalidParentheses) removeInvalidParentheses(word string) []string {
	result := make([]string, 0)
	set := make(map[string]bool)
	queue := make([]string, 0)
	if r.isValid(word) {
		return append(result, word)
	}
	queue = append(queue, word)
	for len(queue) != 0 {
		temp := queue[0]
		queue = queue[1:]
		if r.isValid(temp) {
			result = append(result, temp)
			continue
		}
		if len(result) != 0 {
			continue
		}
		for i := range len(temp) {
			cur := ""
			if temp[i] != '(' && temp[i] != ')' {
				continue
			}
			if i != len(temp)-1 {
				cur = temp[0:i] + temp[i+1:]
			} else {
				cur = temp[0:i]
			}
			if set[cur] == false {
				queue = append(queue, cur)
				set[cur] = true
			}
		}
	}

	return result

}

func (r *RemoveInvalidParentheses) isValid(word string) bool {
	count := 0
	for i := 0; i < len(word); i++ {
		if "(" == word[i:i+1] {
			count++
		} else if ")" == word[i:i+1] {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}
