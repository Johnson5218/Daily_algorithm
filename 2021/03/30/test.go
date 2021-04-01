package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1,2,2,1}
	//arr2 := []int{2,2}
	ok := plusOne(arr1)
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

func intersect(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}

func plusOne(digits []int) []int {
	for i := len(digits)-1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	res := append([]int{1}, digits...)
	return res
}