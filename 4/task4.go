package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(num int, ch <-chan int, wg *sync.WaitGroup) {
	for message := range ch {
		fmt.Printf("Worker №%d text: %d\n", num+1, message)
	}
	wg.Done()
}

func main() {
	var n int
	fmt.Println("input number of workers")
	fmt.Scan(&n)
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(i, ch, &wg)
	}
	i := 0
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	for {
		select {
		case <-sigs:
			close(ch)
			wg.Wait()
			return
		default:
			time.Sleep(time.Second)
			ch <- i
			i++
		}
	}
}
