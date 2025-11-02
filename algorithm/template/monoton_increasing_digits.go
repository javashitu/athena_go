package template

import "strconv"

type MonotonIncreasingDigits struct{}

func (m *MonotonIncreasingDigits) monotonIncreasingDigits(num int) int {
	strNum := strconv.Itoa(num)
	flag := len(strNum)
	arr := []byte(strNum)
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			arr[i-1] = arr[i-1] - 1
			flag = i
		}
	}
	for i := flag; i < len(arr); i++ {
		arr[i] = '9'
	}
	result, _ := strconv.Atoi(string(arr))
	return result
}
