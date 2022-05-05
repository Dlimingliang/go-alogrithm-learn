package linkedList

type SinglyListNode struct {
	val  int
	next *SinglyListNode
}

func NewSinglyListNode(val int) SinglyListNode {
	var node = SinglyListNode{
		val:  val,
		next: nil,
	}
	return node
}

type MyLinkedList struct {
	size *int
	head *SinglyListNode
}

func Constructor() MyLinkedList {
	var list = MyLinkedList{}
	a := 0
	list.size = &a
	return list
}

func (list *MyLinkedList) Get(index int) int {
	if index < 0 || index >= *list.size {
		return -1
	}
	temp := list.head
	for i := 1; i <= index; i++ {
		temp = temp.next
	}
	return temp.val
}

func (list *MyLinkedList) AddAtHead(val int) {
	node := NewSinglyListNode(val)
	node.next = list.head
	list.head = &node
	*list.size++
}

func (list *MyLinkedList) AddAtTail(val int) {
	if *list.size == 0 {
		list.AddAtHead(val)
	} else {
		node := NewSinglyListNode(val)
		temp := list.head
		for i := 1; i < *list.size; i++ {
			temp = temp.next
		}
		temp.next = &node
		*list.size++
	}
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > *list.size {
		return
	}
	if index == 0 {
		list.AddAtHead(val)
	} else if index == *list.size {
		list.AddAtTail(val)
	} else {
		node := NewSinglyListNode(val)
		temp := list.head
		for i := 1; i < index; i++ {
			temp = temp.next
		}
		node.next = temp.next
		temp.next = &node
		*list.size++
	}
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= *list.size {
		return
	}
	if index == 0 {
		list.head = list.head.next
	} else {
		temp := list.head
		for i := 1; i < index; i++ {
			temp = temp.next
		}
		temp.next = temp.next.next
	}
	*list.size--
}
