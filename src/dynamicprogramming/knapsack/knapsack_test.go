package knapsack

import (
	"fmt"
	"testing"
)

func TestKnapsack(t *testing.T) {
	weight := []int{2, 2, 4, 6, 3}
	w := 9
	fmt.Println(knapsack(weight, len(weight), w))
}

func TestKnapsack_enhancement(t *testing.T) {
	weight := []int{1, 1, 5, 1, 2}
	w := 14
	fmt.Println(knapsack_enhancement(weight, w))
}

func TestKanpsackMaxValue(t *testing.T) {
	weight := []int{2, 4, 4, 6, 3}
	values := []int{3, 4, 8, 9, 6}
	w := 9
	fmt.Println(knapsackMaxValue(weight, values, w))
	fmt.Println(knapsackMaxValue_enhance(weight, values, w))
}
