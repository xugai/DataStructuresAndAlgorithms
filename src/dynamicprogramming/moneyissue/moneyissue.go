package moneyissue

var minCountOfCoin = -1
// money表示要付的钱是多少
func minCoins(money int) {
	if money == 1 || money == 3 || money == 5 {
		minCountOfCoin = 1
	}
	states := make([][]bool, money)
	for i := 0; i < money; i++ {
		states[i] = make([]bool, money + 1)
	}
	if money > 5 {
		states[0][1] = true
		states[0][3] = true
		states[0][5] = true
	} else if money > 3 {
		states[0][1] = true
		states[0][3] = true
	} else {
		states[0][1] = true
	}
	for i := 1; i < money; i++ {
		for j := 1; j <= money; j++ {
			// 最快到达money的时候，就是所花费硬币个数最少的时候
			if states[i - 1][money] {
				minCountOfCoin = i
				return
			}
			if states[i - 1][j] {
				if j + 1 <= money && !states[i][j + 1] {
					states[i][j + 1] = true
				}
				if j + 3 <= money && !states[i][j + 3]  {
					states[i][j + 3] = true
				}
				if j + 5 <= money && !states[i][j + 5] {
					states[i][j + 5] = true
				}
			}
		}
	}
}
