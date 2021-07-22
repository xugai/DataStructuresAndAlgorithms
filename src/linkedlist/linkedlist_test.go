package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	p := constructLinkedListWithLoop()
	for true {
		if p.Next == nil {
			t.Log(p.Value)
			break
		}
		t.Log(p.Value)
		p = p.Next
	}
}

func TestLinkedListReverse(t *testing.T) {
	p := constructLinkedListWithLoop()
	head := p
	for true {
		if p.Next == nil {
			fmt.Printf("%s ", p.Value)
			break
		}
		fmt.Printf("%s ", p.Value)
		p = p.Next
	}
	np := linkedListReverse(head)
	fmt.Println()
	for true {
		if np.Next == nil {
			fmt.Printf("%s ", np.Value)
			break
		}
		fmt.Printf("%s ", np.Value)
		np = np.Next
	}
	fmt.Println()
}

func TestIsCycleLinkedList(t *testing.T) {
	//head := constructCycleLinkedList()
	a := Node{"a", nil}
	b := Node{"b", nil}
	a.Next = &b
	b.Next = &a
	t.Log(isCycleLinkedList(&a))
}

func TestDeleteLinkedListFromTail(t *testing.T) {
	head := constructLinkedListWithLoop()
	p := head
	for true {
		if p.Next == nil {
			fmt.Printf("%s ", p.Value)
			break
		}
		fmt.Printf("%s ", p.Value)
		p = p.Next
	}
	head = deleteLinkedListFromTail(head, 5)
	p = head
	fmt.Println()
	for true {
		if p.Next == nil {
			fmt.Printf("%s ", p.Value)
			break
		}
		fmt.Printf("%s ", p.Value)
		p = p.Next
	}
}

func sentinelMode(a *[]int, key int) int {
	if a == nil || len(*a) == 0 {
		return -1
	}
	n := len(*a)
	if (*a)[n - 1] == key {
		return n - 1
	}
	tmp := (*a)[n - 1]
	(*a)[n - 1] = key
	i := 0
	for (*a)[i] != key {
		i++
	}
	(*a)[n - 1] = tmp
	if i == n - 1 {
		return -1
	} else {
		return i
	}
}

func sentinelMode2(a *[]int, key int) int {
	arr := *a
	if a == nil || len(arr) == 0 {
		return -1
	}
	n := len(arr)
	if arr[n - 1] == key {
		return n - 1
	}
	tmp := arr[n - 1]
	arr[n - 1] = key
	i := 0
	for arr[i] != key {
		i++
	}
	arr[n - 1] = tmp
	if i == n - 1 {
		return -1
	} else {
		return i
	}
}

func TestSentinelMode(t *testing.T) {
	a := []int{1, 3, 5, 7, 8}
	key := 9
	t.Log(sentinelMode(&a, key))
}

func BenchmarkSentinelMode(b *testing.B) {
	b.ResetTimer()
	key := 9
	for i := 0; i < b.N; i++ {
		a := []int{}
		for j := 0; j < b.N; j++ {
			a = append(a, j)
		}
		_ = sentinelMode(&a, key)
	}
}

func BenchmarkSentinelMode2(b *testing.B) {
	b.ResetTimer()
	key := 9
	for i := 0; i < b.N; i++ {
		a := []int{}
		for j := 0; j < b.N; j++ {
			a = append(a, j)
		}
		_ = sentinelMode2(&a, key)
	}
}
