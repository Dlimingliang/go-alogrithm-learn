package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(isValid("([)]"))
	//fmt.Println(dailyTemperatures([]int{30, 60, 90}))
	//fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	//fmt.Println(numIslands([][]byte{
	//	{'1', '1', '1', '1', '0'},
	//	{'1', '1', '0', '1', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '0', '0', '0'},
	//}))
	fmt.Println(findTargetSumWays([]int{1}, 1))
}

func findTargetSumWays(nums []int, target int) int {
	var res int
	res = findTargetSumWaysRecursion(nums, 0, 0, target, res)
	return res
}

func findTargetSumWaysRecursion(nums []int, index, sum, target, res int) int {
	if len(nums) == index {
		if target == sum {
			res++
		}
		return res
	}
	res = findTargetSumWaysRecursion(nums, index+1, sum+nums[index], target, res)
	res = findTargetSumWaysRecursion(nums, index+1, sum-nums[index], target, res)
	return res
}

func numIslands(grid [][]byte) int {
	result, n, m := 0, len(grid), len(grid[0])
	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if grid[y][x] == '1' {
				numIslandsRecursion(grid, x, y, n, m)
				result++
			}
		}
	}
	return result
}

func numIslandsRecursion(grid [][]byte, x, y, n, m int) {

	if x < 0 || x == m || y < 0 || y == n || grid[y][x] == '0' {
		return
	}

	grid[y][x] = '0'
	numIslandsRecursion(grid, x+1, y, n, m)
	numIslandsRecursion(grid, x-1, y, n, m)
	numIslandsRecursion(grid, x, y+1, n, m)
	numIslandsRecursion(grid, x, y-1, n, m)
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			tempPre := stack[len(stack)-2]
			tempNext := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if token == "+" {
				stack = append(stack, tempPre+tempNext)
			} else if token == "-" {
				stack = append(stack, tempPre-tempNext)
			} else if token == "*" {
				stack = append(stack, tempPre*tempNext)
			} else if token == "/" {
				stack = append(stack, tempPre/tempNext)
			}
		} else {
			intToken, _ := strconv.Atoi(token)
			stack = append(stack, intToken)
		}
	}
	return stack[0]
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, value := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && temperatures[stack[len(stack)-1]] < value {
				index := stack[len(stack)-1]
				result[index] = i - index
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	return result
}

func isValid(s string) bool {
	bytes := []byte(s)
	stack := make([]byte, 0)
	for _, value := range bytes {
		if value == '(' || value == '[' || value == '{' {
			stack = append(stack, value)
		} else if value == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if value == ']' {
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if value == '}' {
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
