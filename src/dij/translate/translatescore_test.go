package translate

import (
	"fmt"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	mem := make([][]int, 3)
	mem[0] = make([]int, 3)
	mem[1] = make([]int, 4)
	mem[2] = make([]int, 2)
	mem[0][0] = 8
	mem[0][1] = 6
	mem[0][2] = 3
	mem[1][0] = 6
	mem[1][1] = 5
	mem[1][2] = 4
	mem[1][3] = 2
	mem[2][0] = 9
	mem[2][1] = 6
	calculateScore(0, mem, 0)
	k := 3
	for i := 0; i < k; i++ {
		fmt.Printf("%d ", scores.Pop())
	}
}
