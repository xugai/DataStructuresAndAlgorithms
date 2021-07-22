package countingsort

// 计数排序
func CountingSort(arr []int) *[]int {
	if len(arr) <= 1 {
		return &arr
	}
	// 确定桶范围
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	// 构造桶序列
	bucket := make([]int, max + 1)
	// 开始计数
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	// 累加桶内计数值
	for i := 1; i < len(bucket); i++ {
		bucket[i] = bucket[i - 1] + bucket[i]
	}
	// 计数排序
	r := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		index := bucket[arr[i]] - 1
		r[index] = arr[i]
		bucket[arr[i]]--
	}
	//for i := 0; i < len(arr); i++ {
	//	index := bucket[arr[i]] - 1
	//	r[index] = arr[i]
	//	bucket[arr[i]]--
	//}
	// 回写覆盖到原数组中
	for i := 0; i < len(r); i++ {
		arr[i] = r[i]
	}
	return &arr
}
