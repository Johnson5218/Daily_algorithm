package main

import "fmt"

func shellSort(arr []int) {
	length := len(arr)
	for step := length; step >= 1; step /= 2 {
		for i := step; i < length; i += step {
			for j := i - step; j >= 0; j -= step {
				if arr[j+step] < arr[j] {
					swap(arr, j+step, j)
					continue
				}
				break
			}
		}
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}


func main() {
	var arr = []int{1, 2, 2, 5, 7, 89, 4, 3, 1, 4, 5, 6, 79, 1, 3, 4, 5}
	fmt.Println(arr)
	shellSort(arr)
	fmt.Println(arr)
}
