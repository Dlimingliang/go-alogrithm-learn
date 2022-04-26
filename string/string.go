package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(addBinary("11", "1"))
	//fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords3("God Ding"))
}

func reverseWords3(s string) string {
	sl := strings.Fields(s)
	var buffer bytes.Buffer
	for i := 0; i < len(sl); i++ {
		runes := []rune(sl[i])
		for j := len(runes) - 1; j >= 0; j-- {
			buffer.WriteRune(runes[j])
		}
		if i != len(sl)-1 {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
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
