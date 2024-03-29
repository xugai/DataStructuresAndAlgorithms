package heap

import (
	"fmt"
	"testing"
)

func TestHeapify(t *testing.T) {
	points := [][]int{
		{-63, -55},		// 83.6301381082203
		{-20, 17},		// 26.248809496813376
		{-88, -82},		// 120.28299963003916
		{-90, -95},		// 130.86252328302402
		{-88, 18},		// 89.82204629154248
		{-62, -21},		// 65.45991139621256
		{71, -64},		// 95.58765610684259    <---
		{-14, 56},		// 57.723478758647246
		{65, 90},		// 111.01801655587259
		{-48, -52},		// 70.76722405181653
		{59, 92},		// 109.29318368498559
		{-44, -59},		// 73.60027173862879
		{-3, -66},		// 66.06814663663572
	}
	fmt.Println(kClosest(points, 7))
}
