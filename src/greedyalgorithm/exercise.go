package greedyalgorithm

import "strconv"

/*
	在一个非负整数 a 中，我们希望从中移除 k 个数字，让剩下的数字值最小，如何选择移除哪 k 个数字呢？
*/
func remove(num, k int) int {
	numOfString := strconv.Itoa(num)
	n := len(numOfString)
	numOfArr := make([]int, n)
	for i := 0; i < n; i++ {
		val, _ := strconv.Atoi(string(numOfString[i]))
		numOfArr[i] = val
	}
	for k > 0 {
		isEnd := true
		for i := 0; i < n - 1; i++ {
			if numOfArr[i] != -1 && numOfArr[i] > numOfArr[i + 1] {
				numOfArr[i] = -1
				k--
				isEnd = false
				break
			}
		}
		if isEnd {
			for i := n - 1; i >=0; i-- {
				if numOfArr[i] != -1 {
					numOfArr[i] = -1
					k--
					break
				}
			}
		}
	}
	result := ""
	for i := 0; i < n; i++ {
		if numOfArr[i] != -1 {
			result += strconv.Itoa(numOfArr[i])
		}
	}
	valueOfResult, _ := strconv.Atoi(result)
	return valueOfResult
}
