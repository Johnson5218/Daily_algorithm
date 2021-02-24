package main

import "fmt"

func quickSort(arr []int){
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	povid := arr[left]
	for left < right {
		if arr[left+1] > povid {
			arr[left+1], arr[right] = arr[right], arr[left+1]
			right--
		} else if arr[left+1] < arr[left] {
			arr[left+1], arr[left] = arr[left], arr[left+1]
			left++
		} else {
			left++
		}
	}
	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

var arr = []int{1, 2, 2, 5, 7, 89, 4, 3, 1, 4, 5, 6, 79, 1, 3, 4, 5}

func main() {
	quickSort(arr)
	fmt.Println(arr)
}