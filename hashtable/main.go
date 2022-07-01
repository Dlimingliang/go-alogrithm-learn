package main

import (
	"fmt"
	"github.com/Dlimingliang/go-alogrithm-learn/hashtable/hashset"
)

func main() {
	mySet := hashset.Constructor()
	mySet.Add(1)
	mySet.Add(2)
	fmt.Println(mySet.Contains(1))
	fmt.Println(mySet.Contains(3))
	mySet.Add(2)
	fmt.Println(mySet.Contains(2))
	mySet.Remove(2)
	fmt.Println(mySet.Contains(2))
}
