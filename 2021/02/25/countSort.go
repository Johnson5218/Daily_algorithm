package main

import "fmt"

func countSort(arr []int) {
	// 1.找到数组中的最大值和最小值
	min, max := arr[0], arr[0]
	for _, value := range arr {
		if min > value {
			min = value
		}
		if max < value {
			max = value
		}
	}
	// 2.统计数组，每当找到一个数，就将计数数组的相应位置+1
	count := make([]int, max+1, max+1)
	for _, value := range arr {
		count[value]++
	}
	// 3.重新构造数组，将统计数组中的元素按照顺序重新编入原数组
	idx := 0
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			for j := count[i]; j > 0; j-- {
				arr[idx] = i
				idx++
			}
		}
	}
}

func main() {
	var arr = []int{1, 2, 2, 5, 7, 89, 4, 3, 1, 4, 5, 6, 79, 1, 3, 4, 5}
	fmt.Println(arr)
	countSort(arr)
	fmt.Println(arr)
}