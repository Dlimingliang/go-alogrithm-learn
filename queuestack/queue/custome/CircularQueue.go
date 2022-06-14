package my

type CircularQueue struct {
	Queue   []int
	Size    int
	MaxSize int
	Head    int
}

func CircularQueueConstructor(k int) CircularQueue {
	cq := CircularQueue{
		Queue:   make([]int, k, k),
		MaxSize: k,
		Size:    0,
		Head:    0,
	}
	return cq
}

func (this *CircularQueue) EnQueue(value int) bool {
	if !this.IsFull() {
		tail := (this.Head + this.Size) % this.MaxSize
		this.Queue[tail] = value
		this.Size++
		return true
	}
	return false
}

func (this *CircularQueue) DeQueue() bool {
	if !this.IsEmpty() {
		this.Head = (this.Head + 1) % this.MaxSize
		this.Size--
		return true
	}
	return false
}

func (this *CircularQueue) Front() int {
	if !this.IsEmpty() {
		return this.Queue[this.Head]
	}
	return -1
}

func (this *CircularQueue) Rear() int {
	if !this.IsEmpty() {
		tail := (this.Head + this.Size - 1) % this.MaxSize
		return this.Queue[tail]
	}
	return -1
}

func (this *CircularQueue) IsEmpty() bool {
	if this.Size == 0 {
		return true
	}
	return false
}

func (this *CircularQueue) IsFull() bool {
	if this.Size == this.MaxSize {
		return true
	}
	return false
}
