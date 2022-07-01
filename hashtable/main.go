package main

import (
	"fmt"
	"github.com/Dlimingliang/go-alogrithm-learn/hashtable/hashmap"
)

func main() {
	//mySet := hashset.Constructor()
	//mySet.Add(1)
	//mySet.Add(2)
	//fmt.Println(mySet.Contains(1))
	//fmt.Println(mySet.Contains(3))
	//mySet.Add(2)
	//fmt.Println(mySet.Contains(2))
	//mySet.Remove(2)
	//fmt.Println(mySet.Contains(2))

	myMap := hashmap.Constructor()
	myMap.Put(1, 1)
	myMap.Put(2, 2)
	fmt.Println(myMap.Get(1))
	fmt.Println(myMap.Get(3))
	myMap.Put(2, 1)
	fmt.Println(myMap.Get(2))
	myMap.Remove(2)
	fmt.Println(myMap.Get(2))

}
