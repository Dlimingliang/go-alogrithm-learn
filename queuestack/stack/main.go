package main

import "fmt"

func main() {
	//fmt.Println(isValid("([)]"))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
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
