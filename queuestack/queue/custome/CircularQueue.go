package my

type MyCircularQueue struct {
	Queue []int
	Size  int
	Head  int
}

func Constructor(k int) MyCircularQueue {
	cq := MyCircularQueue{
		Queue: make([]int, k, k),
		Size:  0,
		Head:  0,
	}
	return cq
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if !this.IsFull() {
		tail := (this.Head + this.Size) % cap(this.Queue)
		this.Queue[tail] = value
		this.Size++
		return true
	}
	return false
}

func (this *MyCircularQueue) DeQueue() bool {
	if !this.IsEmpty() {
		this.Head = (this.Head + 1) % cap(this.Queue)
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
		tail := (this.Head + this.Size - 1) % cap(this.Queue)
		return this.Queue[tail]
	}
	return -1
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.Size == 0 {
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
