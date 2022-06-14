package my

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyCircularQueue struct {
	Head     *ListNode
	Tail     *ListNode
	Count    int
	Capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Capacity: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if !this.IsFull() {
		node := &ListNode{
			Val: value,
		}
		if this.Count == 0 {
			this.Head, this.Tail = node, node
		} else {
			this.Tail.Next = node
			this.Tail = this.Tail.Next
		}
		this.Count++
		return true
	}
	return false
}

func (this *MyCircularQueue) DeQueue() bool {
	if !this.IsEmpty() {
		this.Head = this.Head.Next
		this.Count--
		return true
	}
	return false
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Head.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Tail.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.Count == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.Count == this.Capacity {
		return true
	}
	return false
}
