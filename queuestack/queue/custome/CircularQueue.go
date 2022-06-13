package my

type MyCircularQueue struct {
	Queue []int
	Size  int
	Head  int
	Tail  int
}

func Constructor(k int) MyCircularQueue {
	cq := MyCircularQueue{
		Queue: make([]int, k, k),
		Size:  0,
		Head:  -1,
		Tail:  -1,
	}
	return cq
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if !this.IsFull() {
		if this.Head == -1 {
			this.Head++
		}
		this.Tail++
		this.Tail = this.Tail % cap(this.Queue)
		this.Queue[this.Tail] = value
		this.Size++
		return true
	}
	return false
}

func (this *MyCircularQueue) DeQueue() bool {
	if !this.IsEmpty() {
		if this.Head == this.Tail {
			this.Head = -1
			this.Tail = -1
		} else {
			this.Head++
			this.Head = this.Head % cap(this.Queue)
		}
		this.Size--
		return true
	}
	return false
}

func (this *MyCircularQueue) Front() int {
	if !this.IsEmpty() {
		return this.Queue[this.Head]
	}
	return -1
}

func (this *MyCircularQueue) Rear() int {
	if !this.IsEmpty() {
		return this.Queue[this.Tail]
	}
	return -1
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.Head == -1 && this.Tail == -1 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.Size == cap(this.Queue) {
		return true
	}
	return false
}