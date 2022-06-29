package custome

type MyStack struct {
	queue    []int
	temQueue []int
}

func NewMyStack() MyStack {
	return MyStack{
		queue:    make([]int, 0),
		temQueue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	for _, value := range this.queue {
		this.temQueue = append(this.temQueue, value)
	}
	this.queue = this.queue[:0]
	this.queue = append(this.queue, x)
	for _, value := range this.temQueue {
		this.queue = append(this.queue, value)
	}
	this.temQueue = this.temQueue[:0]
}

func (this *MyStack) Pop() int {
	value := this.Top()
	this.queue = this.queue[1:]
	return value
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	if len(this.queue) > 0 {
		return false
	}
	return true
}
