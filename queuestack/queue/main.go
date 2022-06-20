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
	//fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
	fmt.Println(numSquares(1))
}

func numSquares(n int) int {
	var maxN int
	for i := 1; i <= n; i++ {
		temp := i * i
		if temp < n {
			maxN = i
		} else if temp == n {
			return 1
		}
	}
	result := 1
	queue := getSquaresValues(maxN)
	for len(queue) > 0 {
		size := len(queue)
		result++
		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]
			if temp == n {
				return result
			}
			for _, value := range getSquaresValues(maxN) {
				tempValue := value + temp
				if tempValue < n {
					queue = append(queue, tempValue)
				} else if tempValue == n {
					return result
				}
			}
		}
	}
	return result
}

func getSquaresValues(n int) []int {
	result := make([]int, 0)
	for i := 1; i <= n; i++ {
		result = append(result, i*i)
	}
	return result
}

func openLock(deadends []string, target string) int {
	const start = "0000"
	if start == target {
		return 0
	}

	deadMap := make(map[string]bool)
	for _, value := range deadends {
		deadMap[value] = true
	}
	if deadMap[start] {
		return -1
	}

	result := 0
	queue := make([]string, 0)
	queue = append(queue, start)
	visited := make(map[string]bool, 0)
	visited[start] = true

	for len(queue) > 0 {
		length := len(queue)
		result++
		for i := 0; i < length; i++ {
			temp := queue[0]
			queue = queue[1:]
			for _, value := range getValue(temp) {
				if !deadMap[value] && !visited[value] {
					if value == target {
						return result
					}
					queue = append(queue, value)
					visited[value] = true
				}
			}
		}
	}
	return -1
}

func getValue(str string) []string {
	result := make([]string, 0)
	bytes := []byte(str)
	for i, value := range bytes {
		bytes[i] = value - 1
		if bytes[i] < '0' {
			bytes[i] = '9'
		}
		result = append(result, string(bytes))
		bytes[i] = value + 1
		if bytes[i] > '9' {
			bytes[i] = '0'
		}
		result = append(result, string(bytes))
		bytes[i] = value
	}
	return result
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
