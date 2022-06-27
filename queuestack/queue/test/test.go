package main

import (
	"fmt"
	my "github.com/Dlimingliang/go-alogrithm-learn/queuestack/queue/custome"
)

func main() {
	//cq := my.Constructor(3)
	//fmt.Println(cq.EnQueue(1))
	//fmt.Println(cq.EnQueue(2))
	//fmt.Println(cq.EnQueue(3))
	//fmt.Println(cq.EnQueue(4))
	//fmt.Println(cq.Rear())
	//fmt.Println(cq.IsFull())
	//fmt.Println(cq.DeQueue())
	//fmt.Println(cq.EnQueue(4))
	//fmt.Println(cq.Rear())
	myQueue := my.NewMyQueue()
	myQueue.Push(1)
	myQueue.Push(2)
	fmt.Println(myQueue.Peek())
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Empty())
}
