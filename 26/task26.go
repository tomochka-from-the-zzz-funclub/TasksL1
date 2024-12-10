package main

import (
	"fmt"
	"strings"
)

func AreStringsLettersUnique(str string) bool {
	data := make(map[rune]bool, len(str))
	// складываем руны в мап, если встречаем такую руну, что она есть в мапе, то все буквы не уникальны
	for _, i := range []rune(strings.ToLower(str)) {
		_, ok := data[i]
		if ok {
			return false
		} else {
			data[i] = true
		}
	}
	return true
}

func main() {
	fmt.Printf("%t\n", AreStringsLettersUnique("abcd"))
	fmt.Printf("%t\n", AreStringsLettersUnique("abCdefAaf"))
	fmt.Printf("%t\n", AreStringsLettersUnique("aabcd"))
	fmt.Printf("%t\n", AreStringsLettersUnique("abcdD"))
}
