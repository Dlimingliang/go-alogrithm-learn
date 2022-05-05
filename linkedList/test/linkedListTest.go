package main

import (
	"fmt"
	"github.com/Dlimingliang/go-alogrithm-learn/linkedList"
)

func main() {
	obj := linkedList.Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	fmt.Println(obj.Get(1))
	obj.DeleteAtIndex(1)
	fmt.Println(obj.Get(1))
}
