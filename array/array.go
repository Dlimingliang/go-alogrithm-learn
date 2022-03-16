package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 0, 1, 1, 1, 1}
	result := findMaxConsecutiveOnes(nums)
	fmt.Println(result)
}

func findMaxConsecutiveOnes(nums []int) int {
	//n为数组的长度
	//时间复杂度 一遍循环 O(n)
	//空间复杂度 3个变量 O(1)
	var maxlength, currentLength int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			currentLength = 0
		} else {
			currentLength++
			if currentLength > maxlength {
				maxlength = currentLength
			}
		}
	}
	return maxlength
}
