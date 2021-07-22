package lcs

import "strconv"

// 求两个字符串的最长公共子串长度
func lcs(a, b string, n, m int) int{
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, m)
	}
	// 初始化第一行,a[0...n]与b[0...0]的最长公共子串长度
	for i := 0; i < n; i++ {
		if string(a[i]) == string(b[0]) {
			states[0][i] = 1
		} else if i > 0 {
			states[0][i] = states[0][i - 1]
		} else {
			states[0][i] = 0
		}
	}
	// 初始化第一列,a[0...0]与b[0...m]的最长公共子串长度
	for j := 0; j < m; j++ {
		if string(a[0]) == string(b[j]) {
			states[j][0] = 1
		} else if j > 0 {
			states[j][0] = states[j - 1][0]
		} else {
			states[j][0] = 0
		}
	}
	// 开始填表
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if string(a[i]) == string(b[j]) {
				states[i][j] = max(states[i - 1][j], states[i][j - 1], states[i - 1][j - 1] + 1)
			} else {
				states[i][j] = max(states[i - 1][j], states[i][j - 1], states[i - 1][j - 1])
			}
		}
	}
	return states[n - 1][m - 1]
}

func max(x, y, z int) int {
	maxValue := x
	if y > maxValue {
		maxValue = y
	}
	if z > maxValue {
		maxValue = z
	}
	return maxValue
}

var maxSequence = ""
// maxSeq(0, "2936517", "2")
func longestIncreaseSubString(k int, numStr string, maxNumSeq string) {
	if k == len(numStr) - 1 {
		if len(maxNumSeq) >= len(maxSequence) {
			maxSequence = maxNumSeq
		}
	}
	for i := k + 1; i < len(numStr); i++ {
		right, _ := strconv.Atoi(string(numStr[i]))
		left, _ := strconv.Atoi(string(numStr[k]))
		if right > left {
			longestIncreaseSubString(i, numStr, maxNumSeq + string(numStr[i]))
		}
	}
}