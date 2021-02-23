package main

import "fmt"

func mergeSort(arr []int) []int {
	// divide 操作
	// 拆分数组至一个元素
	l := len(arr)
	if l == 1 {
		return arr
	}
	mid := l / 2
	mergeLeft := mergeSort(arr[:mid])
	mergeRight := mergeSort(arr[mid:])
	return merge(mergeLeft, mergeRight)
}



func merge(l, r []int) []int {
	// 归并操作
	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0, lLen+rLen)
	lIndex, rIndex := 0, 0
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] <= r[rIndex] {
			res = append(res, l[lIndex])
			lIndex++
		} else {
			res = append(res, r[rIndex])
			rIndex++
		}
	}
	if lIndex < lLen {
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}
	return res
}

func main() {
	arr := []int{1, 2, 2, 5, 7, 89, 4, 3, 1, 4, 5, 6, 79, 1, 3, 4, 5}

	fmt.Println(mergeSort(arr))
}