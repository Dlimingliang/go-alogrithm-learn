package main

import "fmt"

func main() {
	//nums := []int{1, 0, 2, 3, 0, 4, 5, 0}
	//result := findMaxConsecutiveOnes(nums)
	//result := findNumbers(nums)
	//sortedSquares(nums)
	//duplicateZeros(nums)
	//merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	//fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
}

func checkIfExist(arr []int) bool {
	//暴力求解 时间复杂度O(n^2) 空间复杂度O(1)
	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr); j++ {
	//		if i != j && arr[i] == 2 * arr[j] {
	//			return true
	//		}
	//	}
	//}
	//return false

	//先求出每一个的二倍值，然后再进行比较 时间复杂度O(n) 空间复杂度O(n)
	arrMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		arrMap[arr[i]*2] = i
	}

	for i := 0; i < len(arr); i++ {
		if value, ok := arrMap[arr[i]]; ok && value != i {
			return true
		}
	}
	return false
}

func removeDuplicates(nums []int) int {
	//0,0,1,1,1,2,2,3,3,4  0,1,2,3,4
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func removeElement(nums []int, val int) int {
	//3,2,2,3  val=3
	index, last := 0, len(nums)
	for index < last {
		if nums[index] == val {
			nums[index] = nums[last-1]
			last--
		} else {
			index++
		}
	}
	return last
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	capacityIndex := len(nums1)
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[capacityIndex-1] = nums1[m-1]
			m--
		} else {
			nums1[capacityIndex-1] = nums2[n-1]
			n--
		}
		capacityIndex--
	}

	for m > 0 {
		nums1[capacityIndex-1] = nums1[m-1]
		m--
		capacityIndex--
	}

	for n > 0 {
		nums1[capacityIndex-1] = nums2[n-1]
		n--
		capacityIndex--
	}
}

func duplicateZeros(arr []int) {
	//1,0,2,3,0,4,5,0
	//1,0,0,2,3,0,0,4
	length := len(arr)
	var moveNum int
	var lastIndex int
	for i := 0; i < length-moveNum; i++ {
		if arr[i] == 0 && i+moveNum+1 < length {
			moveNum++
			lastIndex = i
		}
	}

	for i := length - moveNum - 1; i >= 0; i-- {
		if arr[i] == 0 && i <= lastIndex && moveNum > 0 {
			arr[i+moveNum] = arr[i]
			arr[i+moveNum-1] = arr[i]
			moveNum -= 1
		} else {
			arr[i+moveNum] = arr[i]
		}
	}
	fmt.Println(arr)
}

func sortedSquares(nums []int) []int {
	//-4,-1,0,3,10
	//时间复杂度O(n)
	//空间复杂度O(1)
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}
	return ans
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
