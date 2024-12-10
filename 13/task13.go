package main

import (
	"fmt"
)

func method1(a, b int) {
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}

func method2(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}

func method3(a, b int) {
	a, b = b, a
	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}

func main() {
	a := 5
	b := 10

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)
	method1(a, b)
	method2(a, b)
	method3(a, b)

}
