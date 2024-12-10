package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	sorted := quickSort(arr)
	fmt.Println(sorted)
}
