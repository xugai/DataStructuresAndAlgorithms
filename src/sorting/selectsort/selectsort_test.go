package selectsort

import (
	"fmt"
	"testing"
)

func TestSelectSorting(t *testing.T) {
	arr := []int{5, 3}
	SelectSorting(&arr)
	fmt.Println(arr)
}
