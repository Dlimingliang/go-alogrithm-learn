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
	//fmt.Println(numIslands(
	//	[][]byte{
	//		{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '0', '1', '1'},
	//		{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0'},
	//		{'1', '0', '1', '1', '1', '0', '0', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	//		{'1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	//		{'1', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	//		{'1', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1'},
	//		{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0', '1', '1', '1', '1'},
	//		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1'},
	//		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	//		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	//		{'0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	//		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	//		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	//		{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1'},
	//		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	//	}))
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
}

func openLock(deadends []string, target string) int {
	if "0000" == target {
		return 0
	}

	deadMap := make(map[string]int)
	for i, value := range deadends {
		deadMap[value] = i
	}
	if _, ok := deadMap["0000"]; ok {
		return 0
	}

	result := 0
	queue := make([]string, 0)
	queue = append(queue, "0000")
	visited := make(map[string]int, 0)
	visited["0000"] = 1

	for len(queue) > 0 {
		length := len(queue)
		result++
		for i := 0; i < length; i++ {
			temp := queue[0]
			queue = queue[1:]
			for _, value := range getValue(temp) {
				_, deadOk := deadMap[value]
				_, visitedOk := visited[value]
				if !deadOk && !visitedOk {
					if value == target {
						return result
					}
					queue = append(queue, value)
					visited[value] = 1
				}
			}
		}
	}
	return -1
}

func pre(x rune) rune {
	if x == '0' {
		return '9'
	} else {
		return x - 1
	}
}

func next(x rune) rune {
	if x == '9' {
		return '0'
	} else {
		return x + 1
	}
}

func getValue(str string) []string {
	slice := make([]string, 0)
	runes := []rune(str)
	for i := 0; i < 4; i++ {
		temp := runes[i]
		runes[i] = pre(temp)
		slice = append(slice, string(runes))
		runes[i] = next(temp)
		slice = append(slice, string(runes))
		runes[i] = temp
	}
	return slice
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
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
