package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Print("input number of seconds:")
	fmt.Scan(&n)
	ch := make(chan int)
	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(time.Second)
		}
	}()
	timer := time.After(time.Duration(n) * time.Second) //в канал таймер отправится сообщенька зерез н секунд
	for {
		select {
		case <-timer:
			return
		default:
			fmt.Printf("Receive message: %d\n", <-ch)
		}
	}
}
