package custome

import "math"

type MinStack struct {
	Stack []int
	Min   int
}

func Constructor() MinStack {
	minStack := MinStack{
		Stack: make([]int, 0),
		Min:   math.MaxInt64,
	}
	return minStack
}

func (this *MinStack) Push(val int) {
	if val < this.Min {
		this.Min = val
	}
	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	if this.Min == this.Stack[len(this.Stack)-1] {
		this.Min = math.MaxInt64
		for i := 0; i < len(this.Stack)-1; i++ {
			if this.Stack[i] < this.Min {
				this.Min = this.Stack[i]
			}
		}
	}
	this.Stack = this.Stack[:len(this.Stack)-1]

}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min
}
