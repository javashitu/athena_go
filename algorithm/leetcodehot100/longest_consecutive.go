package leetcodehot100

type LongestConsecutive struct {
}

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	longest := 0
	for _, num := range nums {
		if numSet[num] {
			numSet[num] = false
			curLongest := 0
			leftNum := num - 1
			for numSet[leftNum] {
				numSet[leftNum] = false
				leftNum--
			}
			leftNum++
			rightNum := num + 1
			for numSet[rightNum] {
				numSet[rightNum] = false
				rightNum++
			}

			rightNum--
			curLongest = rightNum - leftNum + 1
			if longest < curLongest {
				longest = curLongest
			}
		}
	}
	return longest
}
