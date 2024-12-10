package main

import "fmt"

func BinarySearch(arr []int, find int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == find {
			return mid
		}

		// Если элемент больше, то ищем в правой половине, иначе в левой
		if arr[mid] < find {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	find := 5
	result := BinarySearch(arr, find)

	if result != -1 {
		// я возвращаб позиуию, а не индекс
		fmt.Printf("Элемент найден на позиции: %d\n", result+1)
	} else {
		fmt.Println("Элемент не найден")
	}
}
