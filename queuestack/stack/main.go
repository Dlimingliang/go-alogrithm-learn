package main

import "fmt"

func main() {
	fmt.Println(isValid("([)]"))
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
