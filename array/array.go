package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 0, 2, 3, 0, 4, 5, 0}
	//result := findMaxConsecutiveOnes(nums)
	//result := findNumbers(nums)
	//sortedSquares(nums)
	//duplicateZeros(nums)
	//merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	//fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	//fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
	//fmt.Println(replaceElements([]int{400}))
	//moveZeroes([]int{0, 1, 0, 3, 12})
	//fmt.Println(sortArrayByParity([]int{1, 0, 3}))
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))
}

func heightChecker(heights []int) int {
	//heights 5,1,2,3,4
	//excepted 1,2,3,4,5
	//result 5
	//时间复杂度 排序的时间复杂度为O(nlogn) 遍历的时间复杂度为O(n) 所以总的时间复杂度为O(nlogn)
	//空间复杂度 O(1)
	//excepted := make([]int, len(heights))
	//copy(excepted, heights)
	//sort.Ints(excepted)
	//result := 0
	//for i := 0; i < len(heights); i++ {
	//	if excepted[i] != heights[i] {
	//		result++
	//	}
	//}
	//return result

	//计数法
	arr := new([101]int)
	for _, height := range heights {
		arr[height]++
	}

	result := 0
	var value int
	for i, j := 1, 0; i < len(arr); i++ {
		value = arr[i]
		for value > 0 {
			if heights[j] != i {
				result++
			}
			j++
			value--
		}
	}
	return result
}

func sortArrayByParity(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; {
		if nums[i]%2 != 0 {
			if nums[length-1]%2 == 0 {
				nums[length-1], nums[i] = nums[i], nums[length-1]
			}
			length--
		} else {
			i++
		}
	}
	return nums
}

func moveZeroes(nums []int) {
	//0,1,0,3,12
	//1,3,12,0,0
	slow, length := 0, len(nums)
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[slow], nums[i] = nums[i], nums[slow]
			slow++
		}
	}
	fmt.Println(nums)
}

func replaceElements(arr []int) []int {
	//17,18,5,4,6,1
	//18,6,6,6,1,-1
	var current int
	max, length := -1, len(arr)
	for i := length - 1; i >= 0; i-- {
		current = arr[i]
		arr[i] = max
		if current > max {
			max = current
		}
	}
	return arr
}

func validMountainArray(arr []int) bool {
	//0,3,2,1 true
	//length := len(arr)
	//if length  < 3 {
	//	return false
	//}
	//
	//var max, index int
	//for i := 0; i < length; i++ {
	//	if arr[i] > max {
	//		max = arr[i]
	//		index = i
	//	}
	//}
	//if index == length - 1 || index == 0 {
	//	return false
	//}
	//
	//for i := 1; i <= index; i++ {
	//	if arr[i] <= arr[i-1] {
	//		return false
	//	}
	//}
	//for i := index + 1; i < length; i++ {
	//	if arr[i] >= arr[i-1] {
	//		return false
	//	}
	//}

	//一次扫描，比上边的快
	i, n := 0, len(arr)
	//递增扫描
	for ; i+1 < n && arr[i] < arr[i+1]; i++ {
	}
	if i == 0 || i == n-1 {
		return false
	}
	//递减扫描
	for ; i+1 < n && arr[i] > arr[i+1]; i++ {
	}
	return i == n-1
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
	//arrMap := make(map[int]int)
	//for i := 0; i < len(arr); i++ {
	//	arrMap[arr[i]*2] = i
	//}
	//
	//for i := 0; i < len(arr); i++ {
	//	if value, ok := arrMap[arr[i]]; ok && value != i {
	//		return true
	//	}
	//}
	arrMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if value, ok := arrMap[arr[i]*2]; ok && value != i {
			return true
		}
		if arr[i]%2 == 0 {
			if value, ok := arrMap[arr[i]/2]; ok && value != i {
				return true
			}
		}
		arrMap[arr[i]] = i
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
