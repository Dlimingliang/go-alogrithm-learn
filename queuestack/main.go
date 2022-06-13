package main

import "fmt"

func main() {
	slice := make([]int, 0, 5)
	slicePoint := &slice
	fmt.Println("当前长度: ", len(*slicePoint))
	fmt.Println("当前容量: ", cap(slice))
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println("当前长度: ", len(*slicePoint))
	fmt.Println("当前容量: ", cap(slice))
	fmt.Println("index0的值: ", slice[0])
	slice = append(slice[:0], slice[1:]...)
	fmt.Println("当前长度: ", len(*slicePoint))
	fmt.Println("当前容量: ", cap(slice))
	slice = slice[1:]
	fmt.Println("当前长度: ", len(*slicePoint))
	fmt.Println("当前容量: ", cap(slice))

}
