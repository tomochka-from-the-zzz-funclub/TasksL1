package main

import (
	"fmt"
)

func main() {
	val := [...]int{2, 4, 6, 8, 10}
	c := make(chan int, len(val))
	defer close(c)
	for _, i := range val {
		go func(i int) {
			c <- i * i
		}(i)
	}
	sum := 0
	for i := 0; i < len(val); i++ {
		sum += <-c
	}
	fmt.Print(sum, "\n")
}
