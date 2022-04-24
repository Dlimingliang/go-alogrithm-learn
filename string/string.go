package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(addBinary("11", "1"))
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	//时间复杂度O(n)
	//空间复杂度O(n)
	s = strings.TrimSpace(s)
	sl := strings.Fields(s)
	var buffer bytes.Buffer
	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] != " " {
			if i == 0 {
				buffer.WriteString(sl[i])
			} else {
				buffer.WriteString(sl[i])
				buffer.WriteString(" ")
			}
		}
	}
	return buffer.String()
}

func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
