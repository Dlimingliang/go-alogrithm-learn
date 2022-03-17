package main

import (
	"fmt"
)

func main() {
	nums := []int{555, 901, 4827, 1771}
	//result := findMaxConsecutiveOnes(nums)
	result := findNumbers(nums)
	fmt.Println(result)
}

func findNumbers(nums []int) int {
	var result int
	//for _, num := range nums {
	//	s := strconv.Itoa(num)
	//	if len(s) % 2 == 0 {
	//		result ++
	//	}
	//}
	for _, num := range nums {
		var size int
		for num > 0 {
			size++
			num = num / 10
		}
		if size%2 == 0 {
			result++
		}
	}
	return result
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
