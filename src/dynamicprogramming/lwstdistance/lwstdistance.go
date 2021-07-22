package lwstdistance

// 莱文斯坦距离，用来衡量两个字符串的相似程度。莱文斯坦距离越长，说明两个字符串的相似程度越低，反之越高

func lwstdistance(a, b string, n, m int) int {
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, m)
	}
	// 初始化第0行的莱文斯坦距离
	for i := 0; i < n; i++ {
		if string(b[0]) == string(a[i]) {
			states[0][i] = i
		} else if i != 0 {
			states[0][i] = states[0][i - 1] + 1
		} else {
			states[0][i] = 1
		}
	}
	// 初始化第0列的莱文斯坦距离
	for j := 0; j < m; j++ {
		if string(a[0]) == string(b[j]) {
			states[j][0] = j
		} else if j > 0 {
			states[j][0] = states[j - 1][0] + 1
		} else {
			states[j][0] = 1
		}
	}
	// 开始填充状态转移表,按行填充
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if string(a[i]) == string(b[j]) {
				states[i][j] = min(states[i - 1][j] + 1, states[i][j - 1] + 1, states[i - 1][j - 1])
			} else {
				states[i][j] = min(states[i - 1][j] + 1, states[i][j - 1] + 1, states[i - 1][j - 1] + 1)
			}
		}
	}
	return states[n - 1][m - 1]
}

func min(x, y, z int) int {
	minValue := x
	if y < minValue {
		minValue = y
	}
	if z < minValue {
		minValue = z
	}
	return minValue
}