package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// func UsingContext(ctx context.Context, wg *sync.WaitGroup) {
// 	// выводим данные, пока контекст не будет закрыт
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Print("goroutine death\n")
// 			wg.Done()
// 			return
// 		default:
// 			fmt.Print("goroutine with context is alive\n")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	// остановка выполнения горутины с использованием закрытия канала
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	ch := make(chan int)
// 	go UsingChanClose(ch, &wg)
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// 	wg.Wait()

// 	// остановка выполнения горутины с использованием канала закрытия
// 	wg.Add(1)
// 	exit := make(chan interface{})
// 	go UsingChan(exit, &wg)
// 	time.Sleep(1 * time.Second)
// 	exit <- struct{}{}
// 	wg.Wait()

// 	// остановка выполнения горутины с использованием контекстов
// 	ctx, cancel := context.WithCancel(context.Background())
// 	wg.Add(1)
// 	go UsingContext(ctx, &wg)
// 	time.Sleep(5 * time.Second)
// 	cancel()
// 	wg.Wait()
// }

func PrintHello1(ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				wg.Done()
				return // Если канал закрыт, выходим из функции
			}

			time.Sleep(3 * time.Second)
			fmt.Println("Hello1")

		}
	}
}

func PrintHello2(ch <-chan time.Time, wg *sync.WaitGroup) {
	for {
		select {
		case <-ch:
			wg.Done()
			return
		default:
			time.Sleep(3 * time.Second)
			fmt.Println("Hello2")
		}
	}
}

func PrintHello3(flag *bool, wg *sync.WaitGroup) {
	for {
		if *flag {
			time.Sleep(2 * time.Second)
			fmt.Println("Hello3")
		} else {
			wg.Done()
			return
		}
	}
}

func PrintHello4(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Hello4") // Изменено на "Hello4"
		}
	}
}

func main() {

	//  закрытие канала
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go PrintHello1(ch, &wg)
	for i := 0; i < 5; i++ {
		ch <- i

	}
	close(ch)
	wg.Wait()

	///запись в канал
	wg.Add(1)
	//ch2 := make(chan int)
	timer := time.After(10 * time.Second)
	go PrintHello2(timer, &wg)
	wg.Wait()

	//флаг
	wg.Add(1)
	flag := true
	go PrintHello3(&flag, &wg)

	time.Sleep(3 * time.Second)
	flag = false

	wg.Wait()

	wg.Add(1)
	//использовние контекста
	ctx, cancel := context.WithCancel(context.Background())
	go PrintHello4(ctx, &wg)

	// После некоторого времени
	time.Sleep(2 * time.Second)
	cancel() // Сигнализируем горутине о завершении
	wg.Wait()
}
