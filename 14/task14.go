package main

import (
	"fmt"
	"reflect"
)

func determineType1(variable interface{}) {
	// используем оператор switch, чтобы определить динамический тип переменной
	switch variable.(type) {
	case int:
		fmt.Printf("Type int: %d\n", variable)
	case string:
		fmt.Printf("Type string: %s\n", variable)
	case bool:
		fmt.Printf("Type bool: %t\n", variable)
	case chan int:
		fmt.Printf("Type chan int: %v\n", variable)
	default:
		fmt.Printf("Unexpected type: %v\n", variable)
	}
}

func determineType2(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("Type:", t, "value:", i)
}

func determineType3(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
	ch := make(chan int)
	fmt.Println("Metod1:")
	determineType1(interface{}(5))
	determineType1(interface{}("Hello"))
	determineType1(interface{}(true))
	determineType1(interface{}(make(chan int)))
	determineType1(interface{}(7.1))

	fmt.Println("Metod2:")
	determineType2(interface{}(5))
	determineType2(interface{}("Hello"))
	determineType2(interface{}(true))
	determineType2(interface{}(ch))
	determineType2(interface{}(7.1))

	fmt.Println("Metod3:")
	determineType3(interface{}(5))
	determineType3(interface{}("Hello"))
	determineType3(interface{}(true))
	determineType3(interface{}(ch))
	determineType3(interface{}(7.1))
}
