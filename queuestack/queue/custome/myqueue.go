package my

type MyQueue struct {
	stack     []int
	tempStack []int
}

func NewMyQueue() MyQueue {
	return MyQueue{
		stack:     make([]int, 0),
		tempStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	if this.Empty() {
		this.stack = append(this.stack, x)
	} else {
		//转移到临时栈内
		for len(this.stack) > 0 {
			value := this.stack[len(this.stack)-1]
			this.tempStack = append(this.tempStack, value)
			this.stack = this.stack[:len(this.stack)-1]
		}
		//正式放入栈
		this.stack = append(this.stack, x)
		for len(this.tempStack) > 0 {
			value := this.tempStack[len(this.tempStack)-1]
			this.stack = append(this.stack, value)
			this.tempStack = this.tempStack[:len(this.tempStack)-1]
		}
	}
}

func (this *MyQueue) Pop() int {
	value := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return value
}

func (this *MyQueue) Peek() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.stack) > 0 {
		return false
	}
	return true
}
