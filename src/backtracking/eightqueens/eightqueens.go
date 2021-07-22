package eightqueens

import "fmt"

var rows = make([]int, 8)
var count = 1
func cal8Queens(row int) {
	if row == 8 {
		fmt.Printf("第%d种放法: \n", count)
		count++
		print(rows)
		return
	}
	// 循环
	for column := 1; column <= 8; column++ {
		if isOK(row, column) {
			rows[row] = column
			cal8Queens(row + 1)		// 递归
		}
	}
}

// 判断棋子放入当前行的当前列，是否可行
func isOK(row, column int) bool {
	leftup, rightup := column - 1, column + 1
	for i := row - 1; i >= 0; i-- {
		// 判断是否在同一列
		if rows[i] == column {
			return false
		}
		// 判断是否在同一左下对角线
		if leftup > 0 && rows[i] == leftup {
			return false
		}
		// 判断是否在同一右下对角线
		if rightup <= 8 && rows[i] == rightup {
			return false
		}
		leftup--
		rightup++
	}
	return true
}

func print(rows []int) {
	for row := 0; row < 8; row++ {
		for column := 1; column <= 8; column++ {
			if rows[row] == column {
				fmt.Printf("Q ")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Println()
	}
}
