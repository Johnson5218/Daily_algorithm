package main

import "fmt"

func main() {
	arr1 := []int{2,2,1}
	ok := singleNumber(arr1)
	fmt.Println(ok)
}

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[v]++
	}
	return false
}

func singleNumber(nums []int) int {
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v < 2 {
			return k
		}
	}
	return 0
}