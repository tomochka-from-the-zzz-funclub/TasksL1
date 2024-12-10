package main

import (
	"fmt"
)

func deleteElement(arr []int, i int) []int {
	if i < 0 || i > len(arr)-1 {
		fmt.Println("Некорректный индекс")
	}
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("Исходный срез:", numbers)

	indexToRemove := 2

	updatedSlice := deleteElement(numbers, indexToRemove)

	fmt.Println("Обновленный срез:", updatedSlice)
}
