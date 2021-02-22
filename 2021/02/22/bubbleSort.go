package main

import "fmt"

func BubbleSort() {
	arr := []int{1, 2, 2, 5, 7, 89, 4, 3, 1, 4, 5, 6, 79, 1, 3, 4, 5}
	for i := len(arr) - 1; i > 0; i-- {
		swag := true
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swag = false
			}
		}
		if swag {
			break
		}
	}
	fmt.Println(arr)
}
