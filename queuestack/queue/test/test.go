package main

import (
	"fmt"

	"github.com/Dlimingliang/go-alogrithm-learn/queuestack/queue/custome"
)

func main() {
	cq := my.Constructor(2)
	fmt.Println(cq.EnQueue(4))
	fmt.Println(cq.Rear())
	fmt.Println(cq.EnQueue(9))
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.Front())
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.IsEmpty())
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.EnQueue(6))
	fmt.Println(cq.EnQueue(4))
	fmt.Println(cq.Rear())
}
