package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	//fmt.Println(findTargetSumWays([]int{1}, 1))
	//root := TreeNode{Val: 1}
	//two := TreeNode{Val: 2}
	//three := TreeNode{Val: 3}
	//two.Left = &three
	//root.Right = &two
	//fmt.Println(inorderTraversal(&root))
	//fmt.Println(decodeString("3[a2[c]]"))
	//fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 0, 0, 0))
	//fmt.Println(updateMatrix([][]int{{0, 1, 1, 0, 0}, {0, 1, 1, 0, 0}, {0, 1, 0, 0, 1}, {1, 1, 1, 1, 0}, {1, 0, 0, 1, 0}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]int, len(rooms))
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			value := queue[0]
			if visited[value] == 0 {
				visited[value] = 1
				for _, key := range rooms[value] {
					queue = append(queue, key)
				}
			}
			queue = queue[1:]
		}
	}
	for _, value := range visited {
		if value == 0 {
			return false
		}
	}
	return true
}

func updateMatrix(mat [][]int) [][]int {
	n := len(mat)
	m := len(mat[0])
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[i][j] = updateMatrixBFS(mat, i, j, n, m)
		}
	}
	return res
}

func updateMatrixBFS(mat [][]int, row, column, n, m int) int {
	res := 0
	queue := make([]int, 0)
	queue = append(queue, column*n+row)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			value := queue[0]
			y := value % n
			x := value / n
			if mat[y][x] == 0 {
				return res
			}

			if y-1 >= 0 {
				queue = append(queue, x*n+(y-1))
			}
			if y+1 < n {
				queue = append(queue, x*n+(y+1))
			}
			if x-1 >= 0 {
				queue = append(queue, (x-1)*n+y)
			}
			if x+1 < m {
				queue = append(queue, (x+1)*n+y)
			}
			queue = queue[1:]
		}
		res++
	}
	return res
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	visited := make(map[int]bool, 0)
	floodFillRecursion(image, sr, sc, color, image[sr][sc], visited)
	return image
}

func floodFillRecursion(image [][]int, sr, sc, color, target int, visited map[int]bool) {
	if sr < 0 || sr == len(image) || sc < 0 || sc == len(image[0]) || image[sr][sc] != target || visited[sc*len(image)+sr] {
		return
	}
	visited[sc*len(image)+sr] = true
	image[sr][sc] = color
	floodFillRecursion(image, sr-1, sc, color, target, visited)
	floodFillRecursion(image, sr+1, sc, color, target, visited)
	floodFillRecursion(image, sr, sc-1, color, target, visited)
	floodFillRecursion(image, sr, sc+1, color, target, visited)
}

func decodeString(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)

}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}

func inorderTraversal(root *TreeNode) (res []int) {
	//var res []int
	//res = inorderTraversalRecursion(root, res)
	//return res
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

func inorderTraversalRecursion(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = inorderTraversalRecursion(root.Left, res)
	res = append(res, root.Val)
	res = inorderTraversalRecursion(root.Right, res)
	return res
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
