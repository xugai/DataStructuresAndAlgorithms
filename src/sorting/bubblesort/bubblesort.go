package bubblesort

func Bubblesort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	tmp := 0
	flag := false
	for i := len(*arr) - 1; i > 0; i-- {
		for j := 0; j < len(*arr) - 1; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				flag = true
				tmp = (*arr)[j + 1]
				(*arr)[j + 1] = (*arr)[j]
				(*arr)[j] = tmp
			}
		}
		// 如果逆序度为0，则说明当前数组是完全有序的，因此提前退出循环
		if !flag {
			break
		}
	}
}

func BubblesortWithNoFlag(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	tmp := 0
	for i := 0; i < len(*arr) - 1; i++ {
		for j := 0; j < len(*arr) - 1; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				tmp = (*arr)[j + 1]
				(*arr)[j + 1] = (*arr)[j]
				(*arr)[j] = tmp
			}
		}
	}
}
