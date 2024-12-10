package main

import (
	"fmt"
	"strconv"
)

func setIthBit(i int, bit int, number int64) int64 {
	// устанавливаем бит с помощью битовых операций
	if bit == 1 {
		return number | (1 << i)
	} else {
		return number &^ (1 << i)
	}
}

func main() {
	var i, bit int
	var number int64
	fmt.Println("Please input your number: ")
	fmt.Scan(&number)
	fmt.Println("Please input your bit: ")
	fmt.Scan(&bit)
	fmt.Println("Please input your i: ")
	fmt.Scan(&i)
	fmt.Printf("Before: %s\n", strconv.FormatInt(number, 2))
	number = setIthBit(i, bit, number)
	fmt.Printf("Before: %s\n", strconv.FormatInt(number, 2))

}

// Установка бита:
// Если bit равно 1, функция должна установить i-й бит в 1. Это делается с помощью битовой операции OR (|):

// (1 << i) - это операция сдвига влево. Она создает маску, которая имеет 1 только в i-м бите и 0 в остальных. Например, если i равно 3, то (1 << 3) даст 0000...0001000 (в двоичном представлении).
// number | (1 << i) - это операция OR между number и маской. Она устанавливает i-й бит в 1, не изменяя остальные биты.
// Сброс бита:
// Если bit равно 0, функция должна сбросить i-й бит в 0. Это делается с помощью битовой операции AND с отрицанием (&^):

// (1 << i) создает ту же маску, что и раньше.
// number &^ (1 << i) - это операция AND с отрицанием. Она сбрасывает i-й бит в 0. Если i-й бит уже 0, то результат не изменится.
