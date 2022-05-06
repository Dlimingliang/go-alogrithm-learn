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
	size int
	head *SinglyListNode
}

func Constructor() MyLinkedList {
	var list = MyLinkedList{}
	list.size = 0
	list.head = &SinglyListNode{
		val:  0,
		next: nil,
	}
	return list
}

func (list *MyLinkedList) Get(index int) int {
	if index < 0 || index >= list.size {
		return -1
	}
	temp := list.head
	for i := 0; i <= index; i++ {
		temp = temp.next
	}
	return temp.val
}

func (list *MyLinkedList) AddAtHead(val int) {
	list.AddAtIndex(0, val)
}

func (list *MyLinkedList) AddAtTail(val int) {
	list.AddAtIndex(list.size, val)
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > list.size {
		return
	}

	node := NewSinglyListNode(val)
	temp := list.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	node.next = temp.next
	temp.next = &node
	list.size++
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= list.size {
		return
	}

	temp := list.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	temp.next = temp.next.next
	list.size--
}
