package main

import "fmt"

func main() {
	//mySet := hashset.Constructor()
	//mySet.Add(1)
	//mySet.Add(2)
	//fmt.Println(mySet.Contains(1))
	//fmt.Println(mySet.Contains(3))
	//mySet.Add(2)
	//fmt.Println(mySet.Contains(2))
	//mySet.Remove(2)
	//fmt.Println(mySet.Contains(2))
	//myMap := hashmap.Constructor()
	//myMap.Put(1, 1)
	//myMap.Put(2, 2)
	//fmt.Println(myMap.Get(1))
	//fmt.Println(myMap.Get(3))
	//myMap.Put(2, 1)
	//fmt.Println(myMap.Get(2))
	//myMap.Remove(2)
	//fmt.Println(myMap.Get(2))
	//fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	//fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	numMap := make(map[int]bool, 0)
	for {
		value, temp := 0, n
		for temp > 0 {
			num := temp % 10
			num = num * num
			temp = temp / 10
			value += num
		}
		if value == 1 {
			return true
		}
		if numMap[value] {
			return false
		}
		numMap[value] = true
		n = value
	}
	return false
}

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	tempMap := make(map[int]bool, 0)
	numMap := make(map[int]bool)
	if len(nums1) > len(nums2) {
		for _, value := range nums1 {
			numMap[value] = true
		}
		for _, value := range nums2 {
			if numMap[value] {
				tempMap[value] = true
			}
		}
	} else {
		for _, value := range nums2 {
			numMap[value] = true
		}
		for _, value := range nums1 {
			if numMap[value] {
				tempMap[value] = true
			}
		}
	}
	for key, _ := range tempMap {
		res = append(res, key)
	}
	return res
}

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func containsDuplicate(nums []int) bool {
	repeatedMap := make(map[int]bool, 0)
	for _, num := range nums {
		if repeatedMap[num] {
			return true
		} else {
			repeatedMap[num] = true
		}
	}
	return false
}
