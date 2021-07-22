package linkedlist

type Node struct {
	Value string
	Next  *Node
}

func constructCycleLinkedList() *Node{
	a := Node{"a", nil}
	b := Node{"b", nil}
	c := Node{"c", nil}
	d := Node{"d", nil}
	e := Node{"e", nil}
	f := Node{"f", nil}

	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = &f
	//f.Next = &a
	f.Next = nil
	return &a
}

func constructLinkedListWithLoop() *Node {
	head := Node{"a", nil}
	p := &head
	for i := 0; i < 4; i++ {
		currentLetter := []rune(p.Value)[0]
		nextLetterAsciiCode := int(currentLetter) + 1
		nextLetter := string(rune(nextLetterAsciiCode))
		p.Next = &Node{nextLetter, nil}
		p = p.Next
	}
	return &head
}

func linkedListReverse(head *Node) *Node{
	nodeArr := []Node{}
	p := head
	if p == nil {
		return nil
	}
	for true {
		if p.Next == nil {
			nodeArr = append(nodeArr, *p)
			break
		}
		nodeArr = append(nodeArr, *p)
		p = p.Next
	}
	for i := len(nodeArr) - 1; i >=0; i-- {
		if i == 0 {
			nodeArr[i].Next = nil
		} else {
			nodeArr[i].Next = &nodeArr[i - 1]
		}
	}
	return &nodeArr[len(nodeArr) - 1]
}

// 判断是否为环形链表，如果快慢指针能够相遇，则说明它们是在一个环形链表上运动
func isCycleLinkedList(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true  // 快慢指针相遇了，说明是环形链表
		}
	}
	return false
}

// 删除单链表中的倒数第n个结点
func deleteLinkedListFromTail(head *Node, n int) *Node {
	length := 0
	p := head
	for true {
		if p.Next == nil {
			length++
			break
		}
		length++
		p = p.Next
	}
	if n > length || n < 0 || n == 0{
		return head
	}
	var tmpNode *Node
	p = head
	i := 0
	for i <= length {
		// 如果是删除头结点
		if n == length {
			head = p.Next
			p.Next = nil
			break
		}
		if i == length - n {
			tmpNode.Next = p.Next
			break
		}
		tmpNode = p
		p = p.Next
		i++
	}
	return head
}


