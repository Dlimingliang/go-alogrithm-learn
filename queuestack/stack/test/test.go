package main

import (
	"fmt"
	"github.com/Dlimingliang/go-alogrithm-learn/queuestack/stack/custome"
)

func main() {
	//minStack := custome.Constructor()
	//minStack.Push(-2)
	//minStack.Push(0)
	//minStack.Push(-3)
	//fmt.Println(minStack.GetMin()) // 返回 -3.
	//minStack.Pop()
	//fmt.Println(minStack.Top())    // 返回 0.
	//fmt.Println(minStack.GetMin()) // 返回 -2.
	myStack := custome.NewMyStack()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Empty())
}
