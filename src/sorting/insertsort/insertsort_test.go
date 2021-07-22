package insertsort

import (
	"fmt"
	"testing"
)

func TestInsertSorting(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	InsertSorting(&arr)
	fmt.Println(arr)
}
