package knapsack

// 动态规划：空间换时间的一种算法思想

func knapsack(weight []int, n, w int) int {
	states := make([][]bool, n)
	for i := 0; i < n; i++ {
		states[i] = make([]bool, w + 1)
	}
	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}
	for i := 1; i < n; i++ {
		// 当前物品不放入背包
		for k := 0; k <= w; k++ {
			if states[i - 1][k] {
				states[i][k] = true
			}
		}
		// 当前物品放入背包
		for k := 0; k <= w - weight[i]; k++ {
			if states[i - 1][k] {
				states[i][k + weight[i]] = true
			}
		}
	}
	for i := w; i >= 0; i-- {
		if states[n - 1][i] {
			return i
		}
	}
	return -1
}

func knapsack_enhancement(weight []int, w int) int {
	n := len(weight)
	states := make([]bool, w + 1)
	states[0] = true
	if weight[0] <= w {
		states[weight[0]] = true
	}
	for i := 1; i < n; i++ {
		for k := w - weight[i]; k >=0; k-- {
			if states[k] {
				states[k + weight[i]] = true
			}
		}
	}
	for i := w; i >= 0; i-- {
		if states[i] {
			return i
		}
	}
	return -1
}

func knapsackMaxValue(weights, values []int, w int) int {
	n := len(weights)
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, w + 1)
	}
	// 初始化
	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			states[i][j] = -1
		}
	}
	// 特殊处理第一个物品的场景
	states[0][0] = 0
	if weights[0] <= w {
		states[0][weights[0]] = values[0]
	}
	for i := 1; i < n; i++ {
		// 当前物品不放入背包
		for k := 0; k <= w; k++ {
			if states[i - 1][k] != -1 {
				states[i][k] = states[i - 1][k]
			}
		}
		// 当前物品放入背包
		for k := 0; k <= w - weights[i]; k++ {
			if states[i - 1][k] != -1 {
				v := states[i - 1][k] + values[i]
				if v > states[i][k + weights[i]] {
					states[i][k + weights[i]] = v
				}
			}
		}
	}
	maxV := -1
	for i := 0; i <= w; i++ {
		if states[n - 1][i] > maxV {
			maxV = states[n - 1][i]
		}
	}
	return maxV
}

func knapsackMaxValue_enhance(weights, values []int, w int) int{
	n := len(weights)
	states := make([]int, w + 1)
	for i := 0; i <= w; i++ {
		states[i] = -1
	}
	states[0] = 0
	if weights[0] <= w {
		states[weights[0]] = values[0]
	}
	for i := 1; i < n; i++ {
		// 将当前物品放入背包
		for k := w - weights[i]; k >= 0; k-- {
			if states[k] != -1 {
				v := states[k] + values[i]
				if v > states[k + weights[i]] {
					states[k + weights[i]] = v
				}
			}
		}
	}
	maxV := -1
	for i := 0; i <= w; i++ {
		if states[i] > maxV {
			maxV = states[i]
		}
	}
	return maxV
}
