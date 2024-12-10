package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	chan1 := make(chan int)
	chan2 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, i := range arr {
			chan1 <- i
		}
		close(chan1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range chan1 {
			chan2 <- val * 2
		}
		close(chan2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range chan2 {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
