package main

import "fmt"

func insertSort() {
	arr := []int{1, 2, 2, 5, 7, 89, 4, 3, 1, 4, 5, 6, 79, 1, 3, 4, 5}
	for i := 1; i < len(arr); i++ {
		var pos, value int
		value = arr[i]
		pos = i
		for (pos > 0) && (arr[pos-1] > value) {
			arr[pos], arr[pos-1] = arr[pos-1], arr[pos]
			pos--
		}
		arr[pos] = value
	}
	fmt.Println(arr)
}

func main() {
	insertSort()
}
