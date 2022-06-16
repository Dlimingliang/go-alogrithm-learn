package main

import "fmt"

func main() {
	//slice := make([]int, 0, 5)
	//slicePoint := &slice
	//fmt.Println("当前长度: ", len(*slicePoint))
	//fmt.Println("当前容量: ", cap(slice))
	//slice = append(slice, 1)
	//slice = append(slice, 2)
	//fmt.Println("当前长度: ", len(*slicePoint))
	//fmt.Println("当前容量: ", cap(slice))
	//fmt.Println("index0的值: ", slice[0])
	//slice = append(slice[:0], slice[1:]...)
	//fmt.Println("当前长度: ", len(*slicePoint))
	//fmt.Println("当前容量: ", cap(slice))
	//slice = slice[1:]
	//fmt.Println("当前长度: ", len(*slicePoint))
	//fmt.Println("当前容量: ", cap(slice))
	fmt.Println(numIslands(
		[][]byte{
			{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '0', '1', '1'},
			{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0'},
			{'1', '0', '1', '1', '1', '0', '0', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1'},
			{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		}))
}

func numIslands(grid [][]byte) int {
	result, n, m := 0, len(grid), len(grid[0])
	slice := make([]int, 0)
	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if grid[y][x] == '1' {
				slice = append(slice, y+x*n)
				grid[y][x] = '0'
				result++
				for len(slice) > 0 {
					num := len(slice)
					for i := 0; i < num; i++ {
						temp := slice[0]
						slice = slice[1:]
						tempX := temp / n
						tempY := temp % n
						if tempY-1 >= 0 && grid[tempY-1][tempX] == '1' {
							grid[tempY-1][tempX] = '0'
							slice = append(slice, tempY-1+tempX*n)
						}
						if tempY+1 < n && grid[tempY+1][tempX] == '1' {
							grid[tempY+1][tempX] = '0'
							slice = append(slice, tempY+1+tempX*n)
						}
						if tempX-1 >= 0 && grid[tempY][tempX-1] == '1' {
							grid[tempY][tempX-1] = '0'
							slice = append(slice, tempY+(tempX-1)*n)
						}
						if tempX+1 < m && grid[tempY][tempX+1] == '1' {
							grid[tempY][tempX+1] = '0'
							slice = append(slice, tempY+(tempX+1)*n)
						}
					}
				}
			}
		}
	}
	return result
}
