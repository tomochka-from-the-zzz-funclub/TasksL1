package main

import (
	"fmt"
)

func reverseString1(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func reverseString2(str string) string {
	resString := ""
	// Преобразуем строку для корректного итерирования по символам строки в слайс рун
	strArray := []rune(str)
	//обходим слайс рун в обратном порядке, записываем их в результирующую строку
	for i := len(strArray) - 1; i >= 0; i-- {
		resString += string(strArray[i])
	}
	return resString
}

func main() {
	var input string
	fmt.Scan(&input)
	reversed1 := reverseString1(input)
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Перевернутая строка: %s\n", reversed1)
	reversed2 := reverseString2(input)
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Перевернутая строка: %s\n", reversed2)
}
