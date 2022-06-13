package main

import (
	"fmt"

	"github.com/Dlimingliang/go-alogrithm-learn/queuestack/queue/custome"
)

func main() {
	cq := my.Constructor(3)
	fmt.Println(cq.EnQueue(1))
	fmt.Println(cq.EnQueue(2))
	fmt.Println(cq.EnQueue(3))
	fmt.Println(cq.EnQueue(4))
	fmt.Println(cq.Rear())
	fmt.Println(cq.IsFull())
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.EnQueue(4))
	fmt.Println(cq.Rear())
}
