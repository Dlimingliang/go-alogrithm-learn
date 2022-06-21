package custome

import "math"

type MinStack struct {
	Stack []int
	Min   []int
}

func Constructor() MinStack {
	minStack := MinStack{
		Stack: make([]int, 0),
		Min:   []int{math.MaxInt64},
	}
	return minStack
}

func (this *MinStack) Push(val int) {
	min := this.GetMin()
	if val < min {
		this.Min = append(this.Min, val)
	} else {
		this.Min = append(this.Min, min)
	}
	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.Min = this.Min[:len(this.Min)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}
