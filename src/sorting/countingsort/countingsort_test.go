package countingsort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := []int{2, 5, 3, 0, 2, 3, 0, 3}
	result := CountingSort(arr)
	fmt.Println(*result)
}
