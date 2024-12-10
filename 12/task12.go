package main

import "fmt"

func createUniqueStrings1(strings []string) []string {
	var uniqueStrings []string

	contains := func(s []string, str string) bool {
		for _, v := range s {
			if v == str {
				return true
			}
		}
		return false
	}

	for _, str := range strings {
		if !contains(uniqueStrings, str) {
			uniqueStrings = append(uniqueStrings, str)
		}
	}

	return uniqueStrings
}
func createUniqueStrings2(strings []string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, str := range strings {
		set[str] = struct{}{} // Пустая структура для экономии памяти
	}
	return set
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	uniqueStrings1 := createUniqueStrings1(strings)
	uniqueStrings2 := createUniqueStrings2(strings)

	fmt.Println("Уникальные строки способом 1:")
	for _, uniqueStr := range uniqueStrings1 {
		fmt.Println(uniqueStr)
	}
	fmt.Println("Уникальные строки способом 2:")
	for uniqueStr, _ := range uniqueStrings2 {
		fmt.Println(uniqueStr)
	}
}
