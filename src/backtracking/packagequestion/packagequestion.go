package packagequestion

var maxW = -1

/*
	i 表示当前第i个物品
	cw 表示目前为止放入背包的总承重是多少
	w 表示背包能够承受的最大重量
	n 表示一个有多少个物品
	items 表示每个物品的重量是多少
*/
func packagequestion(i, cw, w, n int, items []int) {
	if cw == w || i == n {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	packagequestion(i + 1, cw, w, n, items)		// 不放当前遍历到的物品
	if items[i] + cw <= w {
		packagequestion(i + 1, cw + items[i], w, n, items)		// 放入当前遍历到的物品
	}
}
