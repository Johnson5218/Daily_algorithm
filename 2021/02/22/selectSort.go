package main

import "fmt"

func SelectSort() {
	arr := []int{1, 2, 2, 5, 7, 89, 4, 3, 1, 4, 5, 6, 79, 1, 3, 4, 5}
	var index int
	for i := 0; i < len(arr); i++ {
		index = i
		for j := i; j < len(arr); j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}
		if index != i {
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	fmt.Println(arr)
}

func main() {
	SelectSort()
}
