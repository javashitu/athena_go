package leetcodetop100

type IsAvailable struct{}

func (i *IsAvailable) isAvailable(s string) bool {
	arr := []byte(s)
	stack := make([]byte, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == '(' {
			stack = append(stack, ')')
		} else if arr[i] == '[' {
			stack = append(stack, ']')
		} else if arr[i] == '{' {
			stack = append(stack, '}')
		} else {
			if len(stack) == 0 {
				return false
			}
			str := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if str != arr[i] {
				return false
			}
		}
	}
	return len(stack) == 0
}
